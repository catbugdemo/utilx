package main

import (
	"github.com/pkg/errors"
	"os"
)

// CheckNotExist 判断文件是否存在
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// Mkdir 创建文件夹
func Mkdir(src string) error {
	if err := os.MkdirAll(src, os.ModePerm); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
