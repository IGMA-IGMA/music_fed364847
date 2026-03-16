package main

import (
	"os"
	"path/filepath"
)

func createFile(dir string, filename string) (*os.File, error) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	pathFile := FileJoin(dir, filename)
	_, err = os.Stat(pathFile)
	if err == nil {
		err := os.Remove(pathFile)
		if err != nil {
			return nil, err
		}

	}
	file, err := os.Create(pathFile)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func FileJoin(dir string, filename string) string {
	return filepath.Join(dir, filename)
}
