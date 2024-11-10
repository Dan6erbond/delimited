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

	decoder := delimited.NewDecoder(strings.NewReader(`1|{"start":0}|3|4`), delimited.DecoderWithDelimiter("|"))

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

	if val.A != 1 {
		t.Errorf("expected %d got %d", 1, val.A)
	}

	if val.B == nil {
		t.Errorf("expected %d got <nil>", 0)
	}

	if *val.B != 0 {
		t.Errorf("expected %d got %d", 0, *val.B)
	}

	if val.C != "3" {
		t.Errorf("expected %s got %s", "3", val.C)
	}

	if val.D == nil {
		t.Errorf("expected %s got <nil>", "4")
	}

	if *val.D != "4" {
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

	if val.A != 1 {
		t.Errorf("expected %d got %d", 1, val.A)
	}

	if val.B == nil {
		t.Errorf("expected %d got <nil>", 0)
	}

	if *val.B != 0 {
		t.Errorf("expected %d got %d", 0, val.B)
	}

	if val.C != "4" {
		t.Errorf("expected %s got %s", "4", val.C)
	}

	if val.D == nil {
		t.Errorf("expected %s got <nil>", "3")
	}

	if *val.D != "3" {
		t.Errorf("expected %s got %s", "3", *val.D)
	}
}

func TestUnmarshalIgnore(t *testing.T) {
	var val struct {
		A int
		B *int `delimited:"ignore"`
		C string
		D *string
	}

	delimited.Unmarshal([]byte(`1,3,4`), &val)

	if val.A != 1 {
		t.Errorf("expected %d got %d", 1, val.A)
	}

	if val.B != nil {
		t.Errorf("expected <nil> got %d", val.B)
	}

	if val.C != "3" {
		t.Errorf("expected %s got %s", "3", val.C)
	}

	if val.D == nil {
		t.Errorf("expected %s got <nil>", "4")
	}

	if *val.D != "4" {
		t.Errorf("expected %s got %s", "4", *val.D)
	}
}

func TestUnmarshalUnderflow(t *testing.T) {
	var val struct {
		A int
		B *int `delimited:"ignore"`
		C string
		D *string
	}

	delimited.Unmarshal([]byte(`1,3`), &val)

	if val.A != 1 {
		t.Errorf("expected %d got %d", 1, val.A)
	}

	if val.B != nil {
		t.Errorf("expected <nil> got %d", val.B)
	}

	if val.C != "3" {
		t.Errorf("expected %s got %s", "3", val.C)
	}

	if val.D != nil {
		t.Errorf("expected <nil> got %s", *val.D)
	}
}

func TestUnmarshalOverflow(t *testing.T) {
	var val struct {
		A int
		B *int `delimited:"ignore"`
		C string
		D *string
	}

	delimited.Unmarshal([]byte(`1,3,4,4`), &val)

	if val.A != 1 {
		t.Errorf("expected %d got %d", 1, val.A)
	}

	if val.B != nil {
		t.Errorf("expected <nil> got %d", val.B)
	}

	if val.C != "3" {
		t.Errorf("expected %s got %s", "3", val.C)
	}

	if val.D == nil {
		t.Errorf("expected %s got <nil>", "4")
	}

	if *val.D != "4" {
		t.Errorf("expected %s got %s", "4", *val.D)
	}
}
