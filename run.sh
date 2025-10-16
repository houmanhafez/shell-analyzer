#!/bin/bash

cd "$(dirname "$0")"

go build -o app main.go
./app
