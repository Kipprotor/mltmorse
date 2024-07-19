package main

import (
	"bufio"
	"flag"
	"os"
)

func main() {
	// コマンドライン引数を解析
	filename := flag.String("o", "morse.wav", "Specify a file name")
	flag.Parse()
	//print(filename)

	// 標準入力を inputs に代入する
	sncr := bufio.NewScanner(os.Stdin)
	sncr.Scan()
	inputs := sncr.Text()

	// 生成する音声のパラメータ
	// sample rate: 44100, frequeny: 840, dotduration: 0.125
	synth := NewSynthesizer(44100, 840, 0.125)

	// 音声のデータを生成
	morseTone := synth.ToSineWave(inputs)

	if err := synth.WriteWav(*filename, morseTone); err != nil {
		panic(err)
	}
}
