package main

import (
	"os"
	"testing"
)

func Test_Rainbow(t *testing.T) {
	printc(os.Stdout, RED, "this is red")
	printc(os.Stdout, GREEN, "this is green")
	printc(os.Stdout, BLUE, "this is blue")
}
