package morse

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// Normalize string character by character before converting Morse code
func Normchr(input rune) rune {
	chr := norm.NFD.String(string(input))
	// 日本語の処理
	Kana := normKanaRune([]rune(chr)[0])
	// 変換可能な文字を大文字に変換
	alpha := unicode.ToUpper(Kana)
	return alpha
}

/*
日本語の正規化用ヘルパー関数
入力は1文字のみ
(半)濁点を大文字に、ひらがなならばカタカナに変換する
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
