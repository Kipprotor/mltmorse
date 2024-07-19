package main

import (
  "math"
  "os"

  "github.com/go-audio/audio"
  "github.com/go-audio/wav"
)

type Synthesizer struct {
  sampleRate int
  dot        []int
  dash       []int
  space      []int
}

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

// NewSynthesizer creates a new Synthesizer that can be used to generate sound
func NewSynthesizer(sampleRate int, freq int, dotduration float64) Synthesizer {
  defaultlength := int(dotduration * float64(sampleRate))
  letterspace := make([]int, defaultlength)
  dot := append(genSineWave(freq, sampleRate, dotduration), letterspace...)
  dash := append(genSineWave(freq, sampleRate, 3*dotduration), letterspace...)

  return Synthesizer{
    sampleRate: sampleRate,
    dot:        dot,
    dash:       dash,
    space:      make([]int, 3*defaultlength),
  }
}

// ToSineWave can be used to generate a sound wave from a string
func (synth Synthesizer) ToSineWave(morse string) []int {
  result := make([]int, synth.sampleRate/2)

  for _, c := range morse {
    switch string(c) {
    case ".":
      result = append(result, synth.dot...)
    case "-":
      result = append(result, synth.dash...)
    default:
      result = append(result, synth.space...)
    }
  }
  return result
}

// WriteWav can be used to save the sound wave as a wav file
func (synth Synthesizer) WriteWav(filename string, wavdata []int) error {
  // オーディオバッファを作成
  buf := &audio.IntBuffer{
    Format: &audio.Format{
      SampleRate:  synth.sampleRate,
      NumChannels: 1,
    },
    Data:           wavdata,
    SourceBitDepth: 8,
  }

  // WAVファイルに保存
  outFile, err := os.Create(filename)
  if err != nil {
    panic(err)
  }
  defer outFile.Close()

  encoder := wav.NewEncoder(outFile, synth.sampleRate, 16, 1, 1)
  if err := encoder.Write(buf); err != nil {
    panic(err)
  }

  if err := encoder.Close(); err != nil {
    panic(err)
  }
  return nil
}
