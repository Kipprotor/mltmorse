package mltmorse_test

import (
	"fmt"

	"github.com/Kipprotor/mltmorse"
)

func Example() {
	text := "MORSE IS AWESOME"

	//Convert to morse
	textInMorse := mltmorse.ToMorse(text)
	fmt.Println(textInMorse)

	//Back to text
	backToText := mltmorse.ToText(textInMorse)
	fmt.Println(backToText)
	//Output: -- --- .-. ... .   .. ...   .- .-- . ... --- -- .
	//MORSE IS AWESOME
}
func ExampleRuneToMorse() {
	ch := 'G'
	str := mltmorse.RuneToMorse(ch)

	fmt.Printf("The letter %c converts to: %s", ch, str)
	//Output: The letter G converts to: --.
}
func ExampleRuneToText() {
	str := "--."
	ch := mltmorse.RuneToText(str)

	fmt.Printf("The morse code %s converts to: %c", str, ch)
	//Output: The morse code --. converts to: G
}
func ExampleNormStr() {
	latinText := "HELLO, WORLD!"
	cyrillicText := "ЭТО МОРЗЕ"
	japaneseText := "きょうは リンゴを 2つ たべました"
	for _, text := range []string{latinText, cyrillicText, japaneseText} {
		for _, r := range text {
			fmt.Printf("%v\n", mltmorse.NormStr(string(r)))
			// hello, world!
			// это морзе
			// キヨウハ リンコ゛ヲ 2ツ タヘ゛マシタ
		}
	}
}
