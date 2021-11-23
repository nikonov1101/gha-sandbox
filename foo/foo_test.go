package foo

import (
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New(0)
	if err == nil {
		t.Fatal("error required but got nil")
	}

	_, err = New(-1)
	if err == nil {
		t.Fatal("error required but got nil")
	}

	foo, err := New(42)
	if err != nil {
		t.Fatal("error must be nil")
	}
	if foo.id != 42 {
		t.Fatal("id must be equal to the supplied one")
	}
	s := foo.String()
	expected := "<*foo: 42>"
	if s != expected {
		t.Fatalf(".String() method failed, expected=%s, got=%s", expected, s)
	}
}
