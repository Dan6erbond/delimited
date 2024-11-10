package delimited

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"
)

type Encoder struct {
	delimiter string
	writer    io.Writer
}

// Encode marshals the given struct's fields using the delimiter as a separator.
// For non-string types the `json.Marshal()` method is used to encode values.
// It will ignore any fields with the `delimited:"ignore"` tag and use the `index` tag to determine the field's position in the string if given.
func (e *Encoder) Encode(v any) error {
	vo := reflect.Indirect(reflect.ValueOf(v))
	t := vo.Type()

	fields := getFields(t)

	for i, j := range fields {
		p, err := marshalField(vo.Field(j))

		if err != nil {
			return err
		}

		_, err = e.writer.Write([]byte(p))

		if err != nil {
			return err
		}

		if i != len(fields)-1 {
			_, err = e.writer.Write([]byte(e.delimiter))

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func marshalField(field reflect.Value) (string, error) {
	if field.Type().Kind() == reflect.String {
		return field.Interface().(string), nil
	} else if field.Type() == reflect.PointerTo(reflect.TypeFor[string]()) {
		return *field.Interface().(*string), nil
	} else {
		b, err := json.Marshal(field.Interface())
		return strings.TrimSuffix(strings.TrimPrefix(string(b), "\""), "\""), err
	}
}

type EncoderOpts func(e *Encoder)

// EncoderwithDelimiter sets a custom delimiter for the encoder to use when encoding.
func EncoderWithDelimiter(delimiter string) EncoderOpts {
	return func(e *Encoder) {
		e.delimiter = delimiter
	}
}

func NewEncoder(writer io.Writer, opts ...EncoderOpts) *Encoder {
	e := &Encoder{writer: writer}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

// Marshal provides a convenience layer to encode struct fields to a string using "," as the delimiter.
func Marshal(v any) ([]byte, error) {
	var b bytes.Buffer

	e := NewEncoder(&b, EncoderWithDelimiter(",")).Encode(v)

	if e != nil {
		return nil, e
	}

	return b.Bytes(), nil
}
