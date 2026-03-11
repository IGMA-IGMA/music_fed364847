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


	path_file := file_path(dir, filename)
	_, err = os.Stat(path_file)
	if err == nil {
		err := os.Remove(path_file)
		if err != nil{
			return nil, err
		}
		
	}
	file, err := os.Create(path_file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func file_path(dir string, filename string) string {
	return filepath.Join(dir, filename)
}
