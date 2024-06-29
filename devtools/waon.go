package main

import (
	"fmt"
)

func main() {
	base := 12353
	//base := 12449
	//for i := 0; i < 88; i++ {
	for i := 0; i < 86; i++ {
		//fmt.Println(string(rune(base+i)))
		hira := rune(base + i)
		kana := rune(hira + 96)
		fmt.Printf("Hira %U: %s, Kana %U: %s\n",
			hira, string(hira), kana, string(kana))
	}
}
