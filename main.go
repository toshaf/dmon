package main

import (
	"fmt"
	"os"
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
			next, err = MapDir(dir)
			if err != nil {
				panic(err)
			}
		}

		last = next
	}
}
