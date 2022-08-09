// Ending a file's name with _test.go tells the go test command that this file contains test functions.
// go test
// go test -v

package greetings

import (
	"fmt"
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	fmt.Println(want)
	msg, err := Hello(name)
	fmt.Println(msg)
	fmt.Println(err)
	fmt.Println(want.MatchString(msg))
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("%s") = "%q", %v, want match for "%#q"`, name, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	fmt.Println(msg)
	fmt.Println(err)
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = "%q", %v, want "", error`, msg, err)
	}
}
