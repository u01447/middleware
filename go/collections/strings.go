/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/2 下午4:56
 * Description:
 **/

package collections

import (
	"fmt"
	"unicode/utf8"
)

func MaxNoRepeated(s string) int {
	// 仅支持英文字符
	chNotRepeatIndex := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := chNotRepeatIndex[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		chNotRepeatIndex[ch] = i
	}
	return maxLength
}

func MaxNoRepeatedChn(s string) int {
	chNotRepeatIndex := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := chNotRepeatIndex[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		chNotRepeatIndex[ch] = i
	}
	return maxLength
}

func StrByte(s string) {
	if s == "" {
		s = "yes,我喜欢你！"
	}
	for i, ch := range []byte(s) {
		fmt.Printf("(%d, %X)", i, ch)
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune 其实是将s进行utf-8解码，解码后的字符会转成unicode,然后放入rune中
		fmt.Printf("(%d, %X)", i, ch)
	}
	fmt.Println()
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %X)", i, ch)
	}
}
