package main

import "testing"

func TestUnpackValidWithNumbersCase(t *testing.T) {
	input := "a4bc2d5e10"
	expected := "aaaabccdddddeeeeeeeeee"

	result, err := StringUnpack(input)
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("unpack result is %s, expected %s", result, expected)
	}
}

func TestUnpackValidWithoutNumbersCase(t *testing.T) {
	input := "abcd"
	expected := "abcd"

	result, err := StringUnpack(input)
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("unpack result is %s, expected %s", result, expected)
	}
}

func TestUnpackEmptyCase(t *testing.T) {
	input := ""
	expected := ""

	result, err := StringUnpack(input)
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("unpack result is %s, expected %s", result, expected)
	}
}

func TestInvalidCase(t *testing.T) {
	input := "4a"
	expected := ""

	result, err := StringUnpack(input)
	if err == nil {
		t.Errorf("expected an error for input %s, but got nil", input)
	}

	if result != expected {
		t.Errorf("unpack result is %s, expected %s", result, expected)
	}
}
