#!/bin/bash

echo "Building wasm..."
tinygo build -o build/drawer.wasm -target=wasi ./src/wasm.go
echo "Done building wasm"