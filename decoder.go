package delimited

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"
)

type Decoder struct {
	delimiter string
	reader    io.Reader
}

func (d *Decoder) Decode(v any) error {
	vo := reflect.Indirect(reflect.ValueOf(v))
	t := vo.Type()

	bytes, err := io.ReadAll(d.reader)

	if err != nil {
		return err
	}

	parts := strings.Split(string(bytes), d.delimiter)

	fields := getFields(t)

	for i, p := range parts {
		if i < len(fields) {
			unmarshalField(p, vo.Field(fields[i]))
		}
	}

	return nil
}

func unmarshalField(p string, field reflect.Value) error {
	if field.Type().Kind() == reflect.String {
		field.Set(reflect.ValueOf(p))
	} else if field.Type() == reflect.PointerTo(reflect.TypeFor[string]()) {
		field.Set(reflect.ValueOf(&p))
	} else {
		t := reflect.New(field.Type()).Interface()

		if err := json.Unmarshal([]byte(p), &t); err != nil {
			if err := json.Unmarshal([]byte(`"`+p+`"`), &t); err != nil {
				return err
			}
		}

		field.Set(reflect.Indirect(reflect.ValueOf(t)))
	}

	return nil
}

type DecoderOpts func(e *Decoder)

func DecoderWithDelimiter(delimiter string) DecoderOpts {
	return func(e *Decoder) {
		e.delimiter = delimiter
	}
}

func NewDecoder(reader io.Reader, opts ...DecoderOpts) *Decoder {
	d := &Decoder{reader: reader}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func Unmarshal(data []byte, v any) error {
	return NewDecoder(bytes.NewReader(data), DecoderWithDelimiter(",")).Decode(v)
}
