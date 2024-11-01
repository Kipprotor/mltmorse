mltmorse
========
[![Documentation](https://godoc.org/github.com/Kipprotor/mltmorse?status.svg)](http://godoc.org/github.com/Kipprotor/mltmorse)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kipprotor/mltmorse)](https://goreportcard.com/report/github.com/Kipprotor/mltmorse)

様々な文字をモールス信号への符号化や復号化を行うパッケージ

このリポジトリーは[gSpera/morse](https://github.com/gSpera/morse)からフォークしています。

機能・特徴
=========
- 文字列の符号化や、モールス信号の復号化
- 数種類の文字の符号化 (ラテン文字、ギリシャ文字、キリル文字)
- モールス符号の変換テーブルの指定

サポートしている文字
============
標準でサポートしている変換表は以下の通りです。

- ラテン文字、記号、数字: [ITU-R M.1677-1](https://www.itu.int/dms_pubrec/itu-r/rec/m/R-REC-M.1677-1-200910-I!!PDF-E.pdf)
- ギリシャ文字: [decodemorsecode.com](https://decodemorsecode.com/greek-alphabet/)
- ハングル: SKATS,[decodemorsecode.com](https://www.mykit.com/kor/ele/morse.htm)
  - ハングル文字が適切に正規化できない問題があり、符号化/復号化に失敗します。
- ひらがな・カタカナ:  [日本アマチュア無線連盟](https://www.jarl.org/Japanese/A_Shiryo/A-C_Morse/morse.htm)

また、以下のように変換表を定義し、独自のものを使うことができます。
使いたい文字によっては、モールス信号に変換する前に文字列を正規化する関数を定義する必要があるかもしれません。
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
ツール
=====
CLIツールが[cmd/morsecli](cmd/morsecli)にあります。
モールス信号への符号化や復号化をすることができます。
```bash
$morsecli -s ja > out.morse
テスト
モールス しんごう です。
^C
$morsecli -D -s ja < out.morse
テスト
モールス シンゴウ デス。
```
詳しくは `--help` を参照してください。

コード例
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
