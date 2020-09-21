package filetool

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func DirIsExist(dir string) (bool, error) {
	fi, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if !fi.IsDir() {
		return false, errors.New("this file is not dir")
	}

	return true, nil
}

func CheckDirName(dirName string) error {
	for _, char := range abolishChars {
		if char == "/" {
			continue
		}
		if strings.Contains(dirName, char) {

			return fmt.Errorf("dir name has illegal char:%s", char)
		}
	}
	return nil
}
