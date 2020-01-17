package auth

import "testing"

func TestHello(t *testing.T) {
	want := "TESTERROR"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}