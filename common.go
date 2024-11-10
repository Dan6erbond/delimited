package delimited

import (
	"reflect"
	"strconv"
)

// getFields returns a slice of field indexes taking into account the "index" and "delimited" tags to filter and order fields correctly.
func getFields(t reflect.Type) (fields []int) {
	var index int

	for i := range t.NumField() {
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

	return
}
