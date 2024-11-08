mltmorse
========
[![Documentation](https://godoc.org/github.com/Kipprotor/mltmorse?status.svg)](http://godoc.org/github.com/Kipprotor/mltmorse)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kipprotor/mltmorse)](https://goreportcard.com/report/github.com/Kipprotor/mltmorse)

mltmorse is a library for Morse encoding/decoding and signal sound generation.

This repository is forked from [gSpera/morse](https://github.com/gSpera/morse).

Features
========
- Encoding of strings and decording of Morse code
- Encording of several types of characters (Latin,Greek,Cyrillic)
- Specify Morse code conversion table

Support
=======
The standard supported conversion tables are as follows

- Latin, Symbol, Number: **ITU-R M.1677-1**
- Greek: [decodemorsecode.com](https://decodemorsecode.com/greek-alphabet/)
- Korean: SKATS,[mykit.com](https://www.mykit.com/kor/ele/morse.htm)
  - There is a problem where Hangul characters cannot be properly normalized, thereby causing encoding/decoding to fail.
- Japanese: [The Japen Amateur Radio League ](https://www.jarl.org/Japanese/A_Shiryo/A-C_Morse/morse.htm)

You can also define a conversion table and use your own as shown below. 
Depending on the characters you want to use, you may need to define a function to normalize the string before converting it to Morse code.
```go
var LatinMorse = mltmorse.EncodingMap{
  'A': ".-",
  'B': "-...",
  'C': "-.-.",
  'D': "-..",
  'E': ".",
  ...
}
```

Tool
====
You can find a cli tool in the [cmd/morsecli](cmd/morsecli) directory
This tool can be used for converting to/from morse and generating CW
```bash
$ morsecli -s lt > out.morse
test
this is morse.
^C
$ cat out.morse
- . ... - .-.-- .... .. ...   .. ...   -- --- .-. ... . .-.-.- .-.-
$ morsecli -D < out.morse
TEST
THIS IS MORSE.
```
For more uses look use `--help`

Examples
========
```go
text := "MORSE IS AWESOME"

//Convert to morse
textInMorse := morse.ToMorse(text)
fmt.Println(textInMorse) //-- --- .-. ... .   .. ...   .- .-- . ... --- -- .

//Back to text
backToText := morse.ToText(textInMorse)
fmt.Println(backToText) //MORSE IS AWESOME
```
