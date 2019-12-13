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
wget $1 -O tmp.jpg
echo convert tmp.jpg -filter spline -resize ${size}x${size} -unsharp 0x1 $name
convert tmp.jpg -filter spline -resize ${size}x${size} -quality 80 -unsharp 0x1 $name
convert tmp.jpg -type Grayscale tmp_bw.jpg
convert tmp_bw.jpg -filter spline -resize ${size}x${size} -quality 80 -unsharp 0x1 bw_$name
ls -alht
