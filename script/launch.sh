#!/bin/bash

# Get the directory where this script resides
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR" || exit

# Trap Ctrl+C (SIGINT) and handle it gracefully
trap 'exit 0' SIGINT

# Build the Go program
cd ../src && go build -o ../build/main.exe main.go && cd ..

# Run the program
./build/main.exe
