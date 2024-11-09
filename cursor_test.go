package delimited_test

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/Dan6erbond/delimited"
)

type TaskCursor struct {
	ID      string
	DueDate time.Time
}

func Example() {
	// Create our cursor
	cursor := &TaskCursor{"1234", time.Time{}}

	// Marshal it
	cursorBytes, _ := delimited.Marshal(cursor)

	fmt.Println(string(cursorBytes))

	// Convert to base64 for transport
	b64 := base64.StdEncoding.EncodeToString(cursorBytes)

	fmt.Println(b64)

	// Decode it back into a cursor
	var c TaskCursor

	cursorString, _ := base64.StdEncoding.DecodeString(b64)

	delimited.Unmarshal(cursorString, &c)

	fmt.Printf("%+v\n", c)

	// Output: 1234,0001-01-01T00:00:00Z
	// MTIzNCwwMDAxLTAxLTAxVDAwOjAwOjAwWg==
	// {ID:1234 DueDate:0001-01-01 00:00:00 +0000 UTC}
}
