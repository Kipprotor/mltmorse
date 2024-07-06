package main

import (
	"bufio"
	"flag"
	//"fmt"
	"math"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

// generateSineWave は、指定された周波数、サンプルレート、持続時間でサイン波を生成します
func genSineWave(freq, sampleRate int, duration float64) []int {
	numSamples := int(float64(sampleRate) * duration)
	wave := make([]int, numSamples)
	freqf64 := float64(freq)
	sampleRatef64 := float64(sampleRate)

	for i := 0; i < numSamples; i++ {
		sample := math.Sin(2 * math.Pi * freqf64 * float64(i) / sampleRatef64)
		wave[i] = int(sample * math.MaxInt16) // 16ビット整数にスケール
	}

	return wave
}

func morseToSineWaves(morse string, dotlen float64, sampleRate int) []int {
	spacelen := int(dotlen * float64(sampleRate))
	space := make([]int, spacelen)
	silence := make([]int, 3*spacelen)

	dot := append(genSineWave(840, sampleRate, dotlen), space...)
	dash := append(genSineWave(840, sampleRate, dotlen*3), space...)

	result := make([]int, sampleRate/2)

	for _, c := range morse {
		switch string(c) {
		case ".":
			result = append(result, dot...)
		case "-":
			result = append(result, dash...)
		default:
			result = append(result, silence...)
		}
	}
	return result
}

func main() {
	// コマンドライン引数を解析
	var filename string
	flag.StringVar(&filename, "o", "morse", "Specify a file name")
	flag.Parse()
	//print(filename)

	// 標準入力を inputs に代入する
	sncr := bufio.NewScanner(os.Stdin)
	sncr.Scan()
	inputs := sncr.Text()

	// 生成する音声のパラメータ
	sampleRate := 44100 // サンプルレート (CD品質)
	duration := 0.125   // 持続時間 (秒)

	// 音声のデータを生成
	morseTone := morseToSineWaves(inputs, duration, sampleRate)

	// オーディオバッファを作成
	buf := &audio.IntBuffer{
		Format: &audio.Format{
			SampleRate:  sampleRate,
			NumChannels: 1,
		},
		Data:           morseTone,
		SourceBitDepth: 8,
	}

	// WAVファイルに保存
	outFile, err := os.Create(filename + ".wav")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	encoder := wav.NewEncoder(outFile, sampleRate, 16, 1, 1)
	if err := encoder.Write(buf); err != nil {
		panic(err)
	}

	if err := encoder.Close(); err != nil {
		panic(err)
	}
}
