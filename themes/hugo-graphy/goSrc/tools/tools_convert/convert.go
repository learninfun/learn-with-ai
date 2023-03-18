package tools_convert

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
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

func FileTostring(filePath string) string {
	dat, err := os.ReadFile(filePath)
	check(err)
	return string(dat)
}

func FileNameNoExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func JsonToString(content any) string {
	byteSlice, err := json.Marshal(content)
	check(err)
	return string(byteSlice)
}

func JsonToStringBeauty(content any) string {
	byteSlice, err := json.MarshalIndent(content, "", "  ")
	check(err)
	return string(byteSlice)
}
