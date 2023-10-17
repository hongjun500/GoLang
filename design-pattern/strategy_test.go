// @author hongjun500
// @date 2023/10/17 16:10
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrategy_SaveFile(t *testing.T) {
	s, err := NewStorageStrategy("file")
	assert.NoError(t, err)
	assert.NoError(t, s.Save("no_encrypt_test.md", []byte("## test, no encrypt")))
	s, err = NewStorageStrategy("encrypt_file")
	assert.NoError(t, err)
	assert.NoError(t, s.Save("encrypt_test.md", []byte("## test, encrypt")))
}
