package convert

import (
	"encoding/json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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
