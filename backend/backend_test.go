package backend

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, from backend!Hello, from INIT!"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}