#!/bin/bash

# ファイル名を指定
filename="$1"

# ファイルが存在するかチェック
if [[ ! -f $filename ]]; then
  echo "File not found!"
  exit 1
fi

# ファイルを1文字ずつ読み込み、文字の前後に " を付けて表示
while IFS= read -r -n1 char; do
  echo "\"${char}\": \"\","
done < ${filename}
