package testfy

import (
	"fmt"
	"unicode/utf8"
)

// 表格驱动的测试，使用testify框架

func GetTheMiddleStr(str string) string {
	// 得到字符数，而不是字节数
	wordCnt := utf8.RuneCountInString(str)
	halfPos1 := wordCnt/2
	halfPos2 := wordCnt/2
	// 长度为偶数取两个，如果是奇数取一个
	if wordCnt % 2 == 0 {
		halfPos1 = wordCnt/2 - 1
	}
	s1 := ""
	s2 := ""
	cnt := 0
	for idx, s := range str {
		fmt.Println("start index: ", idx)
		// 这里的idx是rune字符在字符串里的字节位置，不是字符位置，所以不能使用, 需要使用额外的cnt标志
		if cnt == halfPos1 {
			s1 = string(s)
		} else if cnt == halfPos2 {
			s2 = string(s)
		}
		cnt++
	}

	return s1 + s2
}
