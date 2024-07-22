package mltmorse_test

import (
	"bytes"
	"testing"

	"github.com/Kipprotor/mltmorse"
	"github.com/google/go-cmp/cmp"
)

func TestRuneToMorse(t *testing.T) {
	tm := []struct {
		name   string
		input  rune
		output string
	}{
		{"Simple A", 'A', mltmorse.A},
		{"Hardcoded A", 'A', ".-"},
		{"Non supported rune", 'π', ""},
		{"Lowercase", 'a', ".-"},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			get := mltmorse.RuneToMorse(tt.input)
			if get != tt.output {
				t.Errorf("Expected [%s], got: [%s]", tt.output, get)
			}
		})
	}
}

func TestConverter_ToText(t *testing.T) {
	tm := []struct {
		name      string
		converter mltmorse.Converter
		input     string
		output    string
	}{
		{
			"Simple Text",
			mltmorse.DefaultConverter,
			".-.. --- .-. . --",
			"LOREM",
		},
		{
			"Empty String",
			mltmorse.DefaultConverter,
			"",
			"",
		},
		{
			"Trailing Separator with Handler",
			mltmorse.NewConverter(mltmorse.DefaultMorse,
				mltmorse.WithHandler(func(error) string { return "A" }),
				mltmorse.WithTrailingSeparator(true),
			),
			"",
			"A  ",
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			get := tt.converter.ToText(tt.input)
			if get != tt.output {
				t.Errorf("Expected [%s], got: [%s]", tt.output, get)
			}
		})
	}
}

func TestConverter_ToMorse(t *testing.T) {
	tm := []struct {
		name      string
		converter mltmorse.Converter
		input     string
		output    string
		panics    interface{}
	}{
		{
			"Simple Text",
			mltmorse.DefaultConverter,
			"LOREM",
			".-.. --- .-. . --",
			nil,
		},
		{
			"Empty String",
			mltmorse.DefaultConverter,
			"",
			"",
			nil,
		},
		{
			"Character not supported",
			mltmorse.NewConverter(mltmorse.EncodingMap{}, mltmorse.WithHandler(mltmorse.PanicHandler)),
			"A",
			"",
			mltmorse.ErrNoEncoding{Text: "A"},
		},
		{
			"Character not supported - Ignore",
			mltmorse.NewConverter(mltmorse.EncodingMap{}, mltmorse.WithHandler(mltmorse.IgnoreHandler)),
			"A",
			"",
			nil,
		},
		{
			"Trailing Separator with Handler",
			mltmorse.NewConverter(mltmorse.EncodingMap{},
				mltmorse.WithHandler(func(error) string { return "A" }),
				mltmorse.WithTrailingSeparator(true),
			),

			" ",
			"A ",
			nil,
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if tt.panics != err {
					t.Errorf("Expected panic: %v; got: %v", tt.panics, err)
					if err != nil {
						panic(err)
					}
				}
			}()

			get := tt.converter.ToMorse(tt.input)
			if get != tt.output {
				t.Errorf("Expected [%s], got: [%s]", tt.output, get)
			}
		})
	}
}

func TestToText(t *testing.T) {
	tm := []struct {
		name   string
		input  string
		output string
	}{
		{
			"Simple",
			"--..",
			"Z",
		},
		{
			"Empty",
			"",
			"",
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			out := mltmorse.ToText(tt.input)
			if out != tt.output {
				t.Errorf("Expected: %q; got: %q", tt.output, out)
			}
		})
	}
}

func TestHandler(t *testing.T) {
	t.Run("IgnoreHandler", func(t *testing.T) {
		conv := mltmorse.DefaultConverter
		conv.Handling = mltmorse.IgnoreHandler
		out := conv.ToText("--------")
		if out != "" {
			t.Errorf("Expected \"\", got: %q", out)
		}
	})
	t.Run("PanicHandler", func(t *testing.T) {
		defer func() {
			if out := recover(); out == nil {
				t.Error("Expected Panic")
			}
		}()

		conv := mltmorse.DefaultConverter
		conv.Handling = mltmorse.PanicHandler
		conv.ToText("-------")
	})
}

func TestErrors(t *testing.T) {
	t.Run("ErrNoEncoding", func(t *testing.T) {
		err := mltmorse.ErrNoEncoding{Text: "Text"}
		out := err.Error()
		expected := "No encoding for: \"Text\""
		if out != expected {
			t.Errorf("Expected: %q; got: %q", expected, out)
		}
	})
}

func TestConverter_ToMorseWriter(t *testing.T) {
	tm := []struct {
		name   string
		input  string
		output string
	}{
		{
			"Letter",
			"G",
			"--.",
		},
		{
			"Text",
			"TEXT",
			"- . -..- -",
		},
	}

	buffer := bytes.NewBufferString("")
	writer := mltmorse.ToMorseWriter(buffer)
	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			buffer.Reset()
			writer.Write([]byte(tt.input))
			output := buffer.String()
			if output != tt.output {
				t.Errorf("Expected: %q; got: %q", tt.output, output)
			}
		})
	}
}

func TestConverter_ToTextWriter(t *testing.T) {
	tm := []struct {
		name   string
		input  string
		output string
		err    error
		n      int
	}{
		{
			"Letter",
			"--.",
			"G",
			nil,
			3,
		},
		{
			"Text",
			"- . -..- -",
			"TEXT",
			nil,
			10,
		},
	}

	buffer := bytes.NewBufferString("")
	writer := mltmorse.ToTextWriter(buffer)
	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			buffer.Reset()
			n, err := writer.Write([]byte(tt.input))
			if tt.err != err {
				t.Errorf("Expected error: %v; got: %v", tt.err, err)
			}
			if tt.n != n {
				t.Errorf("Expected n: %v; got: %v", tt.n, n)
			}

			output := buffer.String()
			if output != tt.output {
				t.Errorf("Expected: %q; got: %q", tt.output, output)
			}
		})
	}
}

func TestConverter_CharSeparator(t *testing.T) {
	separator := "separator"
	c := mltmorse.NewConverter(
		mltmorse.DefaultMorse,
		mltmorse.WithCharSeparator(separator),
		mltmorse.WithHandler(mltmorse.PanicHandler),
	)
	out := c.CharSeparator()

	if out != separator {
		t.Errorf("Expected: %q; got: %q", separator, out)
	}
}

func TestConverter_EncodingMap(t *testing.T) {
	expectedMap := mltmorse.DefaultMorse

	c := mltmorse.NewConverter(
		expectedMap,
		mltmorse.WithHandler(mltmorse.PanicHandler),
	)
	out := c.EncodingMap()

	for k := range expectedMap {
		if expectedMap[k] != out[k] {
			t.Errorf("Checking: %q: Expected: %q; got: %q", k, expectedMap[k], out[k])
		}
	}
}

func Test_NewConverter(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("Expected")
			}
		}()
		mltmorse.NewConverter(nil)
	})
}

func TestNormStr(t *testing.T) {
	t.Helper()
	tm := []struct {
		name   string
		input  string
		output string
	}{
		{
			"Uppercase",
			"I ate 2 apples",
			"I ATE 2 APPLES",
		},
		{
			"normalize for Japanese",
			"きょうは リンゴを 2つ たべました",
			"キヨウハ リンコ゛ヲ 2ツ タヘ゛マシタ",
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			out := mltmorse.NormStr(tt.input)
			expect := []rune(tt.output)
			if !cmp.Equal(out, expect) {
				t.Errorf("Expected: %q; got: %q", expect, out)
			}
		})
	}
}
