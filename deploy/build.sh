#!/usr/bin/env bash

cd "$(dirname "$0")/../" || exit

go test ./test/
go build -o main .

docker build -t ozone-platform .