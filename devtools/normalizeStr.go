package main

import (
	"fmt"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// normalize string before converting morse code
func normStr(input string) []rune {
	chr := norm.NFD.String(input)

	var result []rune
	for _, r := range chr {
		// 日本語の処理
		kana := normKanaRune(rune(r))
		// 変換可能な文字を大文字に変換
		alpha := unicode.ToUpper(kana)
		result = append(result, alpha)
	}

	return result
}

/*
日本語の正規化用ヘルパー関数
入力は1文字ずつ
(半)濁点を大文字に、ひらがなをカタカナに変換する
それ以外の文字はそのまま返す
*/
func normKanaRune(input rune) rune {
	// 濁点・半濁点が続く場合は、対応する文字に変換
	//println("char: ", string(r))
	switch input {
	case '\u3099': // 濁点 (゛)
		return rune('゛')
	case '\u309A': // 半濁点 (゜)
		return rune('゜')
	default:
		// ひらがなをカタカナに変換
		if unicode.In(input, unicode.Hiragana) {
			return rune(input + 0x60)
		} else {
			return rune(input)
		}
	}
}

func main() {
	// テストケース
	latinText := "Hello, World!"
	cyrillicText := "Привет, мир!"
	japaneseText := "ぱぴぷぺぽ"
	hangulText := "안녕하세요 모스 신호"
	for _, text := range []string{latinText, cyrillicText, japaneseText, hangulText} {
		fmt.Println("Input:", text)

		fmt.Print("Normstr:")
		normalized := normStr(string(text))
		for _, r := range normalized {
			fmt.Print(string(r), " + ")
		}
		/*
			str := norm.NFD.String(text)
			for _, t := range str {
				fmt.Print(string(t))
			}
		*/
		fmt.Println("\n")
	}
}
