package main

import "testing"

func TestReverse(t *testing.T) {
	if reverse(123) != 321 {
		t.FailNow()
	}

	if reverse(-123) != -321 {
		t.FailNow()
	}

	if reverse(120) != 21 {
		t.FailNow()
	}
}
