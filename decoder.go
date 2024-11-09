package delimited

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type Decoder struct {
	Delimiter string
	reader    io.Reader
}

func (d *Decoder) Decode(v any) error {
	vo := reflect.Indirect(reflect.ValueOf(v))
	t := vo.Type()

	bytes, err := io.ReadAll(d.reader)

	if err != nil {
		return err
	}

	parts := strings.Split(string(bytes), d.Delimiter)

	for i, p := range parts {
		if i < t.NumField() {
			field := t.Field(i)

			if field.Tag.Get("index") == "" || field.Tag.Get("index") == strconv.Itoa(i) {
				unmarshalField(p, vo.Field(i))
			}
		}

		for fi := range vo.NumField() {
			field := t.Field(fi)

			if field.Tag.Get("index") == strconv.Itoa(i) {
				unmarshalField(p, vo.Field(fi))
			}
		}
	}

	return nil
}

func unmarshalField(p string, field reflect.Value) {
	if field.Type().Kind() == reflect.String {
		field.Set(reflect.ValueOf(p))
	} else if field.Type() == reflect.PointerTo(reflect.TypeFor[string]()) {
		field.Set(reflect.ValueOf(&p))
	} else {
		t := reflect.New(field.Type()).Interface()
		json.Unmarshal([]byte(p), &t)
		field.Set(reflect.Indirect(reflect.ValueOf(t)))
	}
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{",", reader}
}

func Unmarshal(data []byte, v any) error {
	return NewDecoder(bytes.NewReader(data)).Decode(v)
}
