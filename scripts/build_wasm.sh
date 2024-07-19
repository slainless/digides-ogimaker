#!/bin/bash

echo "Building wasm..."
# tinygo build -o build/drawer.wasm -target=wasi ./src/wasm.go
GOOS=wasip1 GOARCH=wasm go build -o build/drawer.wasm ./src/wasm.go
echo "Done building wasm"