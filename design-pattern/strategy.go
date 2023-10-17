// @author hongjun500
// @date 2023/10/17 16:00
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 策略模式

package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

// StorageStrategy 存储策略
type StorageStrategy interface {
	Save(name string, data []byte) error
}

var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptStorage{},
}

func NewStorageStrategy(name string) (StorageStrategy, error) {
	s, ok := strategys[name]
	if !ok {
		return nil, fmt.Errorf("strategy %s not found", name)
	}
	return s, nil
}

type fileStorage struct {
}

func (f *fileStorage) Save(name string, data []byte) error {
	return os.WriteFile(name, data, os.ModeAppend)
}

type encryptStorage struct {
}

func (e *encryptStorage) Save(name string, data []byte) error {
	// 加密
	encryptData, err := encrypt(data)
	if err != nil {
		return nil
	}
	return os.WriteFile(name, encryptData, os.ModeAppend)
}

// encrypt 加密 data 数据，使用 md5 加密
func encrypt(data []byte) ([]byte, error) {
	n, err := md5.New().Write(data)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%x", n)), nil
}
