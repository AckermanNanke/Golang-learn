package utils

import (
	"errors"
	"os"
)

// 判断存放日志文件夹是否重名或已存在
func PathExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if err == nil {
		if file.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
