package delimited_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Dan6erbond/delimited"
)

func ExampleDecoder() {
	var val struct {
		A int
		B struct {
			Start int
		}
		C string
		D int
	}

	decoder := delimited.NewDecoder(strings.NewReader(`1|{"start":0}|3|4`))
	decoder.Delimiter = "|"

	decoder.Decode(&val)

	fmt.Printf("%+v\n", val)
	// Output: {A:1 B:{Start:0} C:3 D:4}
}

func ExampleUnmarshal() {
	var val struct {
		A int
		B struct {
			Start int
		}
		C string
		D int
	}

	delimited.Unmarshal([]byte(`1,{"start":0},3,4`), &val)

	fmt.Printf("%+v\n", val)
	// Output: {A:1 B:{Start:0} C:3 D:4}
}

func TestUnmarshalPointer(t *testing.T) {
	var val struct {
		A int
		B *int
		C string
		D *string
	}

	delimited.Unmarshal([]byte(`1,0,3,4`), &val)

	if val.B == nil || *val.B != 0 {
		t.Errorf("expected %d got %d", 0, *val.B)
	}

	if val.D == nil || *val.D != "4" {
		t.Errorf("expected %s got %s", "4", *val.D)
	}
}

func TestUnmarshalIndex(t *testing.T) {
	var val struct {
		A int
		B *int
		C string
		D *string `index:"2"`
	}

	delimited.Unmarshal([]byte(`1,0,3,4`), &val)

	if val.B == nil || *val.B != 0 {
		t.Errorf("expected %d got %d", 0, *val.B)
	}

	if val.D == nil || *val.D != "3" {
		t.Errorf("expected %s got %s", "3", *val.D)
	}
}
