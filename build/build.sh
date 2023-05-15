#!/usr/bin/bash

set -xe

FILE_PATH="../src/main.go"

go build -o app $FILE_PATH
