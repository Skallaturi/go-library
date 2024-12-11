package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func AddFrontmatterToDir(path string) {
	filepath.WalkDir(path, AddFrontmatterToFile)
}

func AddFrontmatterToFile(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}
