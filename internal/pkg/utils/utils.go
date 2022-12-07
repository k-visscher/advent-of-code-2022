package utils

import (
	"os"
	"strconv"
)

func ReadFileToString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func MustParseAsInt(str string) int {
	val, err := strconv.Atoi(str)
	CheckForErr(err)

	return val
}

func CheckForErr(e error) {
	if e != nil {
		panic(e)
	}
}
