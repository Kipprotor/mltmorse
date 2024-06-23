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


Tools
=====
CLIツールが[cmd/morsecli](cmd/morsecli)にあります。
モールス信号への符号化や復号化、モールス信号音のwavファイルの生成をすることができます。
```bash
$morsecli -l ja > out.morse
テスト
モールス しんごう です。
^C
$morsecli -D -l ja < out.morse
てすと
もーるす しんごう です。
```
詳しくは `--help` を参照してください。

Examples
========
```go
```