morsetools
==========
**このリポジトリーは現在開発中です**

モールス信号への符号化や復号化、モールス信号音のwavファイルの生成を行うライブラリ。

このリポジトリーは[gSpera/morse](https://github.com/gSpera/morse)からフォークしています。

Features
========
- 文字列の符号化や、モールス信号の復号化
- 数種類の文字の符号化 (ラテン文字、ギリシャ文字、キリル文字)
- モールス符号の変換テーブルの指定
- モールス信号から、モールス信号音を生成し、wavファイルとして出力
- パリス速度の指定

Supports
========
- ラテン文字、記号、数字: **ITU-R M.1677-1**
- ギリシャ文字: [decodemorsecode.com](https://decodemorsecode.com/greek-alphabet/)
- ハングル: SKATS,[decodemorsecode.com](https://www.mykit.com/kor/ele/morse.htm)
- ひらがな・カタカナ:  [日本アマチュア無線連盟](https://www.jarl.org/Japanese/A_Shiryo/A-C_Morse/morse.htm)

Tools
=====
CLIツールが[cmd/morsecli](cmd/morsecli)にあります。
モールス信号への符号化や復号化、モールス信号音のwavファイルの生成をすることができます。
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