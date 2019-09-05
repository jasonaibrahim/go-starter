#!/usr/bin/env bash

cd "$(dirname "$0")/../" || exit

go test ./test/
go build -o main ./src/

docker build -t go-starter .