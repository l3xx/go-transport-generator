#!/bin/sh

cat os.map | awk '{split($0,part,"|"); print part[1]"\n"part[2]}' | xargs -n 2 bash -c 'GOOS=$0 GOARCH=$1 CGO_ENABLED=0 go build -o ./out/gtg-$0-$1 ./cmd/gtg/main.go'
