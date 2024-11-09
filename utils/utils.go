package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

var watchedFiles = make(map[string]time.Time)

func GetPath() (string, error) {

	ex, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	dir := filepath.Dir(ex)

	err = scanFile(dir)
	if err != nil {

		fmt.Println("error while scanning the file")
		return "", err
	}

	return dir, nil
}

func scanFile(dir string) error {

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if filepath.Ext(path) == ".go" {

			watchedFiles[path] = info.ModTime()
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
		return err
	}

	return nil
}

func CheckForChanges(dir string) (bool, error) {

	changed := false

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if filepath.Ext(path) == ".go" {

			modTime, ok := watchedFiles[path]
			if !ok || info.ModTime().After(modTime) {
				watchedFiles[path] = info.ModTime()
				changed = true
			}
		}
		return nil
	})
	return changed, err
}
