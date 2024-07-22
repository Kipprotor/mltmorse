#!/usr/bin/env bash
echo "Please type what alpabet you want to convert to morse code
lt:Latin, gr:Greek, cy:Cyrillic, kr:Korean, ja:Katakana"
read -p ">" alpha

echo "Alphabet: ${alpha}; Please type what you want to convert into Morse code"
go run ./cmd/morsecli/ -s $alpha