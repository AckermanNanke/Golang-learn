package utils

import (
	"encoding/json"
	"errors"
	"gorm-demo/global"
	"io"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func RequeatHandle(d *interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, d)
	if err != nil {
		global.GVA_LOG.Error("JSON格式转换失败", zap.Error(err))
		return nil, err
	}
	return d, nil
}

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
