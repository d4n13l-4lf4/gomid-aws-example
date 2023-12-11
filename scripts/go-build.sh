#!/bin/bash

# Parse arguments
while [ $# -gt 0 ]
do
key="$1"
case $key in
    --cmd-dir)
    CMD_DIR="$2"
    shift # past argument
    shift
    ;;
    --cmd-dir=*)
    CMD_DIR="${key#*=}"
    shift
    ;;
    --out-dir)
    OUT_DIR="$2"
    shift # past argument
    shift
    ;;
    --out-dir=*)
    OUT_DIR="${key#*=}"
    shift
    ;;
esac
done
CMD_DIR="${CMD_DIR:-cmd}"
OUT_DIR="${OUT_DIR:-build}"

# Erase output directory
rm -rf $PWD/$OUT_DIR

# Compile handlers
find "$PWD/$CMD_DIR" -type 'f' -name '*.go' -print0 | while read -d $'\0' file
do
    echo Compiling $file
    DIR_NAME=$(basename $(dirname $file))
    go build -tags lambda.norpc -ldflags="-s -w" -o $OUT_DIR/$DIR_NAME/bootstrap $file
    zip -r -j "$OUT_DIR/$DIR_NAME.zip" $OUT_DIR/$DIR_NAME
    rm -r $OUT_DIR/$DIR_NAME
    echo $file compiled!
done
echo "Done compiling!"