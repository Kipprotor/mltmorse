package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Kipprotor/mltmorse"
)

func main() {
	var (
		decode    bool
		newlineCode string
		alphabet  string
	)
	flag.BoolVar(&decode, "D", false, "Decodes input (Morse -> Text)")
	flag.StringVar(&newlineCode, "nl", "", "Specify the signal to be used as a newline code. For example -..-.. (nl)")
	flag.StringVar(&alphabet, "s", "lt", "alphabet to use (lt:Latin, gr:Greek, cy:Cyrillic, kr:Korean, ja:Katakana)")

	in := PathFlag("-")
	var out string
	flag.Var(&in, "in", "The input file; default to stdin")
	flag.StringVar(&out, "out", "-", "The output file; default to stdout")
	flag.Parse()

	var writer io.Writer
	if out == "-" {
		writer = os.Stdout
	} else {
		var err error
		writer, err = os.OpenFile(out, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot create output stream:", err)
			os.Exit(1)
		}
	}

	reader, err := in.Stream()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot create reader stream:", err)
		os.Exit(1)
	}

	var DictMap = map[string]mltmorse.EncodingMap{
		"lt": mltmorse.LatinMorse,
		"gr": mltmorse.GreekMorse,
		"cy": mltmorse.CyillicMorse,
		"kr": mltmorse.KoreanMorse,
		"ja": mltmorse.KataMorse,
	}

	encodingMap := mltmorse.MergeEncMap(DictMap[alphabet], mltmorse.NumMorse, mltmorse.SymbolMorse)
	encodingMap['\n'] = newlineCode

	converter := mltmorse.NewConverter(encodingMap,
		mltmorse.WithLowercaseHandling(true),
		mltmorse.WithHandler(mltmorse.IgnoreHandler),
	)

	var stream io.Writer
	if decode {
		stream = converter.ToTextWriter(writer)
	} else {
		stream = converter.ToMorseWriter(writer)
	}

	if _, err := io.Copy(stream, reader); err != nil {
		panic(err)
	}
}
