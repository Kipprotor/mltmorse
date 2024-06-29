package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"
)

var msg = `hira2kana [-s swap|kana|hira] <strings>
-s swap: ひらがなとカタカナを入れ替え
-s hira: 全てのカタカナをひらがなに
-s kata: 全てのひらがなをカタカナに
`

func main() {
	// コマンドライン引数を解析
	swapMode := flag.String("s", "swap", msg)
	flag.Parse()

	args := flag.Args()
	for _, arg := range args {
		//println("args: ",arg)
		fmt.Println(convert(arg, *swapMode))
	}
}

// convert は文字列を指定されたモードで変換します
func convert(input string, mode string) string {
	var result strings.Builder
	for _, r := range input {
		switch mode {
		case "swap":
			switch {
			case unicode.In(r, unicode.Hiragana):
				result.WriteRune(r + 0x60) // ひらがなをカタカナに変換
			case unicode.In(r, unicode.Katakana):
				result.WriteRune(r - 0x60) // カタカナをひらがなに変換
			default:
				result.WriteRune(r) // その他の文字はそのまま
			}
		case "hira":
			if unicode.In(r, unicode.Katakana) {
				result.WriteRune(r - 0x60) // カタカナをひらがなに変換
			} else {
				result.WriteRune(r) // その他の文字はそのまま
			}
		case "kata":
			if unicode.In(r, unicode.Hiragana) {
				result.WriteRune(r + 0x60) // ひらがなをカタカナに変換
			} else {
				result.WriteRune(r) // その他の文字はそのまま
			}
		default:
			//fmt.Printf("Warnning: %v is invalid argument\n", mode)
			result.WriteRune(r) // 不明なモードの場合は変換しない
		}
	}
	return result.String()
}
