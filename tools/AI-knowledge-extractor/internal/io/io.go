package io

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FileCopy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// exists returns whether the given file or directory exists
func FileExists(path string) bool {
	_, error := os.Stat(path)
	// check if error is "file not exists"
	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}

func StringToFile(filePath, content string) {
	fo, err := os.Create(filePath)
	check(err)

	fo.WriteString(content)

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
}

func FileToString(filePath string) string {
	dat, err := os.ReadFile(filePath)
	check(err)
	return string(dat)
}

func FileNameNoExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
