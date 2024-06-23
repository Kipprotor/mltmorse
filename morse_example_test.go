package morse_test

import (
	"fmt"

	"github.com/Kipprtor/morsetools"
)

func Example() {
	text := "MORSE IS AWESOME"

	//Convert to morse
	textInMorse := morse.ToMorse(text)
	fmt.Println(textInMorse)

	//Back to text
	backToText := morse.ToText(textInMorse)
	fmt.Println(backToText)
	//Output: -- --- .-. ... .   .. ...   .- .-- . ... --- -- .
	//MORSE IS AWESOME
}
func ExampleRuneToMorse() {
	ch := 'G'
	str := morse.RuneToMorse(ch)

	fmt.Printf("The letter %c converts to: %s", ch, str)
	//Output: The letter G converts to: --.
}
func ExampleRuneToText() {
	str := "--."
	ch := morse.RuneToText(str)

	fmt.Printf("The morse code %s converts to: %c", str, ch)
	//Output: The morse code --. converts to: G
}
func ExampleNormalizeStr() {
	latinText := "Hello, World!"
	cyrillicText := "Привет, мир!"
	japaneseText := "ぱぴぷぺぽ"
	hangulText := "안녕하세요 세계"
	for _, text := range []string{latinText, cyrillicText, japaneseText, hangulText} {
		for _, r := range text {
			fmt.Println(Normchr(latinText))    // hello, world!
			fmt.Println(Normchr(cyrillicText)) // привет, мир!
			fmt.Println(Normchr(japaneseText)) // ハ゜ヒ゜フ゜ヘ゜ホ゜
			for _, i := range Normchr(hangulText) {
				fmt.Print(string(i), " ")
			}
		}
	}
}
