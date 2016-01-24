package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: %s <dir> <cmd> [args ...]\n", os.Args[0])
		os.Exit(1)
	}

	dir := os.Args[1]

	last, err := MapDir(dir)
	if err != nil {
		panic(err)
	}

	for {
		<-time.After(time.Second)

		next, err := MapDir(dir)
		if err != nil {
			panic(err)
		}

		if !next.Equal(last) {
			Exec(os.Args[2], os.Args[3:])
		}

		last = next
	}
}

type DirMap map[string]time.Time

func MapDir(dir string) (DirMap, error) {
	dm := make(DirMap)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		dm[path] = info.ModTime()

		return nil
	})

	return dm, err
}

func (d1 DirMap) Equal(d2 DirMap) bool {
	if len(d1) != len(d2) {
		return false
	}

	for p1, t1 := range d1 {
		t2, ok := d2[p1]
		if !ok || t1 != t2 {
			return false
		}
	}

	return true
}

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

type Colour string

const (
	RESET Colour = "\u001B[0m"
	RED          = "\u001B[37;41m"
	GREEN        = "\u001B[37;42m"
	BLUE         = "\u001B[37;44m"
)

func printc(w io.Writer, c Colour, s string, args ...interface{}) {
	fmt.Fprint(w, c)
	fmt.Fprint(w, "\n ")
	fmt.Fprintf(w, s, args...)
	fmt.Fprint(w, "\n")
	fmt.Fprint(w, RESET)
	fmt.Fprint(w, "\n")
}
