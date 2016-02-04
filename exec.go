package main

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func Exec(c string, args []string) {
	printc(os.Stderr, BLUE, "%s %s", c, strings.Join(args, " "))
	cmd := exec.Command(c, args...)

	so, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	se, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go copy(os.Stdout, so)
	go copy(os.Stderr, se)

	err = cmd.Wait()
	if err != nil {
		printc(os.Stderr, RED, "O_o %s", err)
	} else {
		printc(os.Stderr, GREEN, ":D")
	}
}

func copy(w io.Writer, r io.Reader) {
	buffer := make([]byte, 1024)

	for {
		nr, err := r.Read(buffer)
		if nr == 0 {
			return
		}
		nw := 0
		for nw < nr {
			i, err := w.Write(buffer[nw:nr])
			if err != nil {
				panic(err)
			}
			nw += i
		}
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
	}
}
