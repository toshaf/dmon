package main

import (
	"fmt"
	"io"
)

type Colour string

const (
	RESET Colour = "\u001B[0m"
	RED          = "\u001B[37;41m"
	GREEN        = "\u001B[37;42m"
	BLUE         = "\u001B[37;44m"
)

func printc(w io.Writer, c Colour, s string, args ...interface{}) {
	fmt.Fprint(w, c)
	fmt.Fprint(w, " \n ")
	fmt.Fprintf(w, s, args...)
	fmt.Fprint(w, " \n ")
	fmt.Fprint(w, RESET)
	fmt.Fprint(w, "\n")
}
