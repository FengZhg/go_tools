package utils

import (
	"os"
	"path/filepath"
)

// @Author: Feng
// @Date: 2022/3/26 18:53

//GetExeFileName 获取运行文件名
func GetExeFileName() (exePath string) {
	defer func() {
		if exePath == "" {
			exePath = "default"
		}
	}()
	//获取可执行文件路径
	path, err := os.Executable()
	if err != nil {
		return
	}
	//分解可执行文件路径
	_, exePath = filepath.Split(path)
	return
}
