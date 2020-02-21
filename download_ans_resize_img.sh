#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "No url"
    exit
fi

if [ -z "$2" ]
then
    echo "No size"
    exit
fi

if [ -z "$3" ]
then
    echo "No name"
    exit
fi

url=$1
size=$2
name=$3

base_name=${name%%.*}
extension=${name##*.}
tmp_file=${base_name}_tmp.$extension
bw_file=${base_name}_bw.$extension

#echo "tmp_file=$tmp_file"
#echo "bw_file=$bw_file"
wget $1 -O $tmp_file
echo convert tmp.jpg -filter spline -resize ${size}x${size} -unsharp 0x1 $name
convert $tmp_file -filter spline -resize ${size}x${size} -quality 80 -unsharp 0x1 $name
convert $tmp_file -type Grayscale $bw_file
convert $bw_file -filter spline -resize ${size}x${size} -quality 80 -unsharp 0x1 $bw_file
ls -alht
