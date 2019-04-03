#!/bin/sh

set -e

GOOS=linux GOARCH=amd64 go build -o build/linux_amd64/encryption  -ldflags '-s -w' main.go
GOOS=windows GOARCH=amd64 go build -o build/windows_amd64/encryption.exe  -ldflags '-s -w' main.go
GOOS=darwin GOARCH=amd64 go build -o build/darwin_amd64/encryption  -ldflags '-s -w' main.go

cd build

tar zcvf linux_amd64.tar.gz linux_amd64
tar zcvf windows_amd64.tar.gz windows_amd64
tar zcvf darwin_amd64.tar.gz darwin_amd64
