package utils

import (
	"advent-of-code-2022/assets"
	"strconv"
)

func MustReadFileToString(path string) string {
	data, err := assets.FS.ReadFile(path)
	CheckForErr(err)

	return string(data)
}

func ReadFileToString(path string) (string, error) {
	data, err := assets.FS.ReadFile(path)
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Min(x int, y int) int {
	if x <= y {
		return x
	}

	return y
}

func Max(x int, y int) int {
	if x >= y {
		return x
	}

	return y
}

func CheckForErr(e error) {
	if e != nil {
		panic(e)
	}
}
