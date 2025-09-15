package example

import "testing"

func TestFoo_Greet(t *testing.T) {
	foo := NewFoo("Tester")
	got := foo.Greet()
	want := "Hello, Tester"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
