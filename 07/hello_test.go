package hello_test

import (
	"testing"
	"hello"
)

func TestReturnGreeting(t *testing.T) {
	want := "Hi there yourself!"
	got := hello.ReturnGreeting("Hi there")
	if want != got {
		t.Errorf("want %q, but got %q", want, got)
	}
}