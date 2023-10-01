#!/usr/bin/env bash

PACKAGE=$1

if [ -z "$PACKAGE" ]; then
    echo "Usage: $0 <package>"
    exit 1
fi

DIST_DIR=build/release

if [ -d $DIST_DIR ]; then
    rm -rf $DIST_DIR
fi

mkdir -p $DIST_DIR

#declare -a platforms=("darwin amd64" "darwin arm64" "linux amd64" "linux arm" "linux arm64" "windows amd64" "windows arm64")
declare -a platforms=("linux amd64" "linux arm" "linux arm64" "windows amd64" "windows arm64")

for platform in "${platforms[@]}"
do
    goos_goarch=($platform)
    goos=${goos_goarch[0]}
    goarch=${goos_goarch[1]}
    exe="${PACKAGE}_${goos}_${goarch}"

    if [ $goos = "windows" ]; then
       exe+=".exe"
    fi

    dist=$DIST_DIR/$exe
    echo "Building $dist"
    env CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -ldflags="-s -w" -o $dist .
    if [ $? -ne 0 ]; then
        echo "Error building ${dist}"
        exit 1
    fi
done
