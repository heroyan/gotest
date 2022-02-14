package testfy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTheMiddleStr(t *testing.T) {
	// 表格驱动测试
	assert := assert.New(t)
	testData := []struct{
		input string
		expected string
	}{
		{"hahaha", "ha"},
		{"hahhaha", "h"},
		{"大中国", "中"},
		{"我们的大中国", "的大"},
		{"", ""},
		{"a", "a"},
		{"aa", "aa"},
		{"中", "中"},
		{"中国", "中国"},
	}
	for _, data := range testData {
		real := GetTheMiddleStr(data.input)
		assert.Equal(data.expected, real)
	}
}
