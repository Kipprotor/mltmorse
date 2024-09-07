package mltmorse_test

import (
	"fmt"

	"github.com/Kipprotor/mltmorse"
)

func Example() {
	texts := []string{"Latin", "ελληνικά", "Русский", "にほんご"}
	maps := []mltmorse.EncodingMap{mltmorse.LatinMorse, mltmorse.GreekMorse, mltmorse.CyillicMorse, mltmorse.KataMorse}
	for i := 0; i < len(texts); i++ {
		conv := mltmorse.NewConverter(maps[i])

		textInMorse := conv.ToMorse(texts[i])
		backToText := conv.ToText(textInMorse)

		fmt.Printf("%v -> %v -> %v\n", texts[i], textInMorse, backToText)
		// Output:
		// Latin -> .-.. .- - .. -. -> LATIN
		// ελληνικά -> . .-.. .-.. .... -. .. -.- .- -> ΕΛΛΗΝΙΚΑ
		// Русский -> .-. ..- ... ... -.- .. .. -> РУССКИИ
		// にほんご -> -.-. -.. .-.-. ---- .. -> ニホンコ゛
	}
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
func ExampleStrNorm() {
	latinText := "hello, world!"
	cyrillicText := "это морзе"
	japaneseText := "きょうは リンゴを 2つ たべました"
	for _, text := range []string{latinText, cyrillicText, japaneseText} {
		r := mltmorse.NormStr(string(text))
		fmt.Printf("%v\n", string(r))
		// Output:
		// HELLO, WORLD!
		// ЭТО МОРЗЕ
		// キヨウハ リンコ゛ヲ 2ツ タヘ゛マシタ
	}
}
