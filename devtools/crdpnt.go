package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputs []string

func main() {
	if len(os.Args[1:]) > 0 {
		// コマンドライン引数があれば、そちらを処理する
		inputs = os.Args[1:]
	} else {
		// コマンドライン引数がない場合は、標準入力から文字列を受け取る
		scnr := bufio.NewScanner(os.Stdin)
		scnr.Scan()
		inputs = strings.Split(scnr.Text(), " ")
	}

	for _, text := range inputs {
		textRune := []rune(text)
		fmt.Printf("text = \"%v\", len = %v\n",
			text, len(textRune))

		fmt.Println("ch| Hex | Dec")
		for _, unc := range textRune {
			//fmt.Printf("char:%c, Unicode: %U, point: %d\n",
			fmt.Printf("%c, %U, %d\n",
				unc, unc, unc)
		}
		fmt.Println()
	}
}
