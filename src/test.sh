#!/bin/bash
  # https://drive.google.com/file/d/1iJhtVwfU2oYY1UQC97CaBBvAdsIRP76A/view?usp=sharing
fileid="1iJhtVwfU2oYY1UQC97CaBBvAdsIRP76A"
filename="manifest2.json"
curl -c ./cookie -s -L "https://drive.google.com/uc?export=download&id=${fileid}" > /dev/null
curl -Lb ./cookie "https://drive.google.com/uc?export=download&confirm=`awk '/download/ {print $NF}' ./cookie`&id=${fileid}" -o ${filename}