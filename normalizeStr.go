package mltmorse

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// Normalize string before converting Morse code; If the input is upper case, it will be converted to upper case. If the input is Japanese, it will be converted to Katakana.
func NormStr(input string) []rune {
	ch := norm.NFD.String(input)

	var result []rune
	for _, r := range ch {
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
	var ch rune
	if unicode.In(input, unicode.Hiragana) {
		ch = input + 0x60
	} else {
		ch = input
	}

	if v, exist := smallKana[ch]; exist {
		return v
	} else {
		return ch
	}
}

var smallKana = map[rune]rune{
	'゙': '゛',
	'゚': '゜',
	'ァ': 'ア',
	'ィ': 'イ',
	'ゥ': 'ウ',
	'ェ': 'エ',
	'ォ': 'オ',
	'ヵ': 'カ',
	'ヶ': 'ケ',
	'ッ': 'ツ',
	'ャ': 'ヤ',
	'ュ': 'ユ',
	'ョ': 'ヨ',
}
