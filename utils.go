package morse

import (
	"io"
)

// mergeMap merges 2 Map objects map[rune]interface{}
func mergeEncMap(m1, m2 EncodingMap) EncodingMap {
	ans := make(EncodingMap, len(m1)+len(m2))

	for _, m := range []EncodingMap{m1, m2} {
		for k, v := range m {
			ans[k] = v
		}
		/*
			for _, c := range m {
				for k, v := range c {
					ans[k] = v
				}
			}
		*/
	}
	return ans
}

// ToText converts a morse string to his textual representation, it is an alias to DefaultConverter.ToText
func ToText(morse string) string { return DefaultConverter.ToText(morse) }

// ToMorse converts a text to his morse rrpresentation, it is an alias to DefaultConverter.ToMorse
func ToMorse(text string) string { return DefaultConverter.ToMorse(text) }

// ToMorseWriter translates all the text written to the returned io.Writer in morse code and writes it in the input io.Writer
func ToMorseWriter(output io.Writer) io.Writer { return DefaultConverter.ToMorseWriter(output) }

// ToTextWriter translates all the text written to the returned io.Writer from morse code and writes it in the input io.Writer
func ToTextWriter(output io.Writer) io.Writer { return DefaultConverter.ToTextWriter(output) }

type translateToMorse struct {
	conv   Converter
	buffer []byte

	input  io.Reader
	output io.Writer
}

// Text -> Morse
func (t translateToMorse) Write(data []byte) (int, error) {
	morse := t.conv.ToMorse(string(data))
	_, err := t.output.Write([]byte(morse))
	return len(data), err
}

type translateToText struct {
	conv   Converter
	buffer []byte

	input  io.Reader
	output io.Writer
}

// Morse -> Text
func (t translateToText) Write(data []byte) (int, error) {
	morse := t.conv.ToText(string(data))
	_, err := t.output.Write([]byte(morse))
	return len(data), err
}

func reverseEncodingMap(encoding EncodingMap) map[string]rune {
	ret := make(map[string]rune, len(encoding))

	for k, v := range encoding {
		ret[v] = k
	}

	return ret
}
