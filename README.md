morsetools
=====
**This repository is UNDER deveplopment**

morsetools is a library for Morse encoding/decoding and signal sound generation.

This repository is forked from [gSpera/morse](https://github.comg/gSpera/morse).

Features
========
- Encoding of strings and decording of Morse code
- Encording of several types of characters (Latin,Greek,Cyrillic)
- Specify Morse code conversion table
- Generate Morse code sound from Morse code and output as wav file
- Specify Paris speed

Support
=======
- Latin, Symbol, Number: **ITU-R M.1677-1**
- Greek: [decodemorsecode.com](https://decodemorsecode.com/greek-alphabet/)
- Korean: SKATS,[mykit.com](https://www.mykit.com/kor/ele/morse.htm)
- Japanese: [The Japen Amateur Radio League ](https://www.jarl.org/Japanese/A_Shiryo/A-C_Morse/morse.htm)
you can use custom ones freely using a custom [EncodingMap](https://github.com/Kipprotor/morsetools/maps.go)

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
