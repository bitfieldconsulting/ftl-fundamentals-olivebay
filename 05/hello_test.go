package hello_test

import (
	"testing"
	"hello"
)

func TestHello(t *testing.T) {
	want := "Hello Gophers!"
	got := hello.Greeting()
	if want != got {
		t.Errorf("want %q, but got %q", want, got)
	}
}
