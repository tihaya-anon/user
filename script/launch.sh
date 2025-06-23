#!/bin/bash
# Navigate to source directory
cd ../src || exit

# Create build directory if it doesn't exist
mkdir -p ../build

# Build the application
go build -o ../build/go_build_MVC_DI.exe main.go

# Run the application
../build/go_build_MVC_DI.exe