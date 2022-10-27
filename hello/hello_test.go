package main

import (
	"testing"
	// "./main"
)

// TestmatchIntString calls matchInt with a non-integer string, should FAIL
func TestMatchIntString(t *testing.T) {
	want := false
	resp := matchInt("hello")
	if want != resp {
		t.Errorf(`matchInt("hello") = %t, should be %t, nil`, resp, want)
	}
}
/*
// TestmatchIntStringInt calls matchInt with an integer string, should PASS
func TestMatchIntStringInt(t *testing.T) {
	want := true
	resp := matchInt("123")
	if want != resp {
		t.Errorf(`matchInt("123") = %t, should be %t`, resp, want)
	}
}

// TestmatchIntEmpty calls matchInt with an empty string, should PASS
func TestMatchIntEmpty(t *testing.T) {
	want := true
	resp := matchInt("")
	if want != resp {
		t.Errorf(`matchInt("") = %t, should be %t`, resp, want)
	}
}
*/