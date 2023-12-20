package FileHandler

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist){
		return false
	}
	return true
}

func ReadFile(fillename string) (string, error) {
	data, err := ioutil.ReadFile(fillename)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(data), "\n\t"), nil
}
