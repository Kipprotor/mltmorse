package main

//package morse

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// normalize string before converting morse code
func Normstr(input string, charT string) string {
	var result strings.Builder

	for _, r := range input {
		switch charT {
		/*
		   case "Hangul":
		     t := norm.NFD.String(string(r))
		     for _, rr := range t {
		       println(string(rr))
		       result.WriteRune(rr)
		     }
		*/
		case "HiraKana": // unicode.In(r, unicode.Hiragana, unicode.Katakana):
			// 日本語を正規化
			result.WriteString(normKana(string(r)))
		default: //"Latin", "Greek", "Cyrillic" の場合を想定
			// 文字を小文字に変換
			result.WriteRune(unicode.ToUpper(r))
		}
	}

	return result.String()
}

/*
日本語の正規化用ヘルパー関数
*/
func normKana(input string) string {
	var result strings.Builder
	nfdStr := norm.NFD.String(input) // NFD正規化で文字を分解
	for _, r := range nfdStr {
		// 濁点・半濁点が続く場合は、対応する文字に変換
		//println("char: ", string(r))
		switch r {
		case '\u3099': // 濁点 (゛)
			result.WriteRune('゛')
		case '\u309A': // 半濁点 (゜)
			result.WriteRune('゜')
		default:
			// ひらがなをカタカナに変換
			if unicode.In(r, unicode.Hiragana) {
				result.WriteRune(r + 0x60)
			} else {
				result.WriteRune(r)
			}
		}
	}
	return result.String()
}

func main() {
	// テストケース
	latinText := "Hello, World!"
	cyrillicText := "Привет, мир!"
	japaneseText := "ぱぴぷぺぽ"
	//hangulText := "안녕하세요 세계"

	fmt.Println(Normstr(latinText, ""))            // hello, world!
	fmt.Println(Normstr(cyrillicText, ""))         // привет, мир!
	fmt.Println(Normstr(japaneseText, "HiraKana")) // は゜ひ゜ふ゜へ゜ほ゜
	//fmt.Println(Normstr(hangulText,"Hangul"))
}
