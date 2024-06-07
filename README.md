morsetools
=====
morsetools is a library for Morse encoding/decoding and signal sound generation.

This repository is forked from [gSpera/morse](https://github.comg/gSpera/morse).

Features
========
- Encoding of strings and decording of Morse code
- Encording of several types of characters
- Specify Morse code conversion table
- Generate Morse code sound from Morse code and output as wav file
- Specify Paris speed

Support
=======
This library supports the default morse (as defined by **ITU-R M.1677-1**) code but custom ones can be used freely using a custom [EncodingMap](https://godoc.org/github.com/gSpera/morse#EncodingMap)

Tool
====
You can find a cli tool in the [cmd/morsecli](cmd/morsecli) directory
This tool can be used for converting to/from morse and generating CW
```bash
$ morse > out.morse
test
this is morse.
^C
$ cat out.morse
- . ... - .-.-- .... .. ...   .. ...   -- --- .-. ... . .-.-.- .-.-
$ morse -D < out.morse
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
