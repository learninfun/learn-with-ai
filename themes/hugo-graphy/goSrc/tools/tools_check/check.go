package tools_check

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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
