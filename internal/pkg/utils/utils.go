package utils

import "os"

func ReadFileToString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func CheckForErr(e error) {
	if e != nil {
		panic(e)
	}
}
