package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func Test_main(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{
			name:    "正常系",
			input:   "hoge fuga",
			counts:  map[rune]int{'f': 1, 'g': 2, 'e': 1, 'a': 1, 'h': 1, 'o': 1, 'u': 1, ' ': 1},
			utflen:  [utf8.UTFMax + 1]int{0, 9, 0, 0, 0},
			invalid: 0,
		},
		{
			name:    "日本語",
			input:   "こんにちは！",
			counts:  map[rune]int{'こ': 1, 'ん': 1, 'に': 1, 'ち': 1, 'は': 1, '！': 1},
			utflen:  [utf8.UTFMax + 1]int{0, 0, 0, 6, 0},
			invalid: 0,
		},
		{
			name:    "絵文字",
			input:   "👍",
			counts:  map[rune]int{'👍': 1},
			utflen:  [utf8.UTFMax + 1]int{0, 0, 0, 0, 1},
			invalid: 0,
		},
		{
			name:    "特殊文字(3バイト)",
			input:   "é",
			counts:  map[rune]int{'é': 1},
			utflen:  [utf8.UTFMax + 1]int{0, 0, 1, 0, 0},
			invalid: 0,
		},
		{
			name:    "空文字",
			input:   "",
			counts:  map[rune]int{},
			utflen:  [utf8.UTFMax + 1]int{0, 0, 0, 0, 0},
			invalid: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			r := bufio.NewReader(strings.NewReader(c.input))
			actualCounts, actualUtflen, actualInvalid := CharCount(r)
			if !reflect.DeepEqual(actualCounts, c.counts) {
				t.Errorf("err in counts. actual: %v, expected: %v", actualCounts, c.counts)
			}
			if !reflect.DeepEqual(actualUtflen, c.utflen) {
				t.Errorf("err in utflen. actual: %v, expected: %v", actualUtflen, c.utflen)
			}
			if actualInvalid != c.invalid {
				t.Errorf("err in invalid. actual: %v, expected: %v", actualInvalid, c.invalid)
			}
		})
	}
}
