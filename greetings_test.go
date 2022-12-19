package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value
func TestHelloName(t *testing.T) {
	name := "Kiran"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Kiran")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Kiran") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for any error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("")= %q, %v, want "", error`, msg, err)
	}
}
