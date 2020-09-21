package filetool

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

func FileIsExist(fileName string) (bool, error) {
	fi, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if fi.IsDir() {
		return false, errors.New("this file is dir")
	}

	return true, nil
}

var abolishChars = []string{"\\", "/", ":", "*", "?", "\"", "<", ">", "|", "^"}

func CheckFileName(fileName string) error {
	_, basef := path.Split(fileName)
	for _, char := range abolishChars {
		if strings.Contains(basef, char) {
			return fmt.Errorf("file name has illegal char:%s", char)
		}
	}
	return nil
}
