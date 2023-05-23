package main

import "testing"

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./shortURLz", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
