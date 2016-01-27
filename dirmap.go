package main

import (
	"os"
	"path/filepath"
	"time"
)

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
