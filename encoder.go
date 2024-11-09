package delimited

import (
	"encoding/json"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type Encoder struct {
	Delimiter string
	writer    io.Writer
}

func (e *Encoder) Encode(v any) error {
	var parts []string

	vo := reflect.Indirect(reflect.ValueOf(v))
	t := vo.Type()

	var fields []int
	index := 0

	for i := range vo.NumField() {
		for j := range t.NumField() {
			if t.Field(j).Tag.Get("index") == strconv.Itoa(i) {
				fields = append(fields, j)
				index++
			}
		}

		if t.Field(i).Tag.Get("delimited") == "ignore" {
			continue
		}

		fields = append(fields, i)
		index++
	}

	for _, j := range fields {
		p, err := marshalField(vo.Field(j))

		if err != nil {
			return err
		}

		parts = append(parts, p)
	}

	e.writer.Write([]byte(strings.Join(parts, e.Delimiter)))
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

func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{",", writer}
}

func Marshal(v any) ([]byte, error) {
	var b strings.Builder

	e := NewEncoder(&b).Encode(v)

	if e != nil {
		return nil, e
	}

	return []byte(b.String()), nil
}
