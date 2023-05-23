package main

import "testing"

func Example_Main() {
	goMain([]string{})
	// Output:
	// Hello World
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./shortURLz", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
