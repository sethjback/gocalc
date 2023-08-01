package main

import "testing"

func TestAdd(t *testing.T) {
	addr := &Adder{}

	addr.num1 = 1
	addr.num2 = 2

	if addr.Calculate() != 3 {
		t.Error("1 and 2 should be 3")
	}
}
