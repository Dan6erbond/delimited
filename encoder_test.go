package delimited_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Dan6erbond/delimited"
)

func ExampleEncoder() {
	var val struct {
		A int
		B struct {
			Start int
		}
		C string
		D int
	} = struct {
		A int
		B struct{ Start int }
		C string
		D int
	}{
		A: 1,
		B: struct{ Start int }{Start: 2},
		C: "3",
		D: 4,
	}

	var b strings.Builder

	enc := delimited.NewEncoder(&b)
	enc.Delimiter = "|"

	enc.Encode(val)

	fmt.Println(b.String())
	// Output: 1|{"Start":2}|3|4
}

func ExampleMarshal() {
	var val struct {
		A int
		B struct {
			Start int
		}
		C string
		D int
	} = struct {
		A int
		B struct{ Start int }
		C string
		D int
	}{
		A: 1,
		B: struct{ Start int }{Start: 2},
		C: "3",
		D: 4,
	}

	b, _ := delimited.Marshal(val)

	fmt.Println(string(b))
	// Output: 1,{"Start":2},3,4
}

func TestMarshalIndex(t *testing.T) {
	var val struct {
		A int
		B struct {
			Start int
		}
		C string `index:"1"`
		D int
	} = struct {
		A int
		B struct{ Start int }
		C string `index:"1"`
		D int
	}{
		A: 1,
		B: struct{ Start int }{Start: 2},
		C: "3",
		D: 4,
	}

	b, _ := delimited.Marshal(val)

	fmt.Println(string(b))
	// Output: 1,3,{"Start":2},4
}

func TestMarshalIgnore(t *testing.T) {
	var val struct {
		A int
		B struct {
			Start int
		}
		C string `delimited:"ignore"`
		D int
	} = struct {
		A int
		B struct{ Start int }
		C string `delimited:"ignore"`
		D int
	}{
		A: 1,
		B: struct{ Start int }{Start: 2},
		C: "3",
		D: 4,
	}

	b, _ := delimited.Marshal(val)

	fmt.Println(string(b))
	// Output: 1,{"Start":2},4
}
