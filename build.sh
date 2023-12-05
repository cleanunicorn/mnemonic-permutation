#!/bin/bash

# A simple build script for a Go-based CLI application.

set -e  # Exit on error

# echo "Running tests..."
# go test ./...

echo "Building the project..."
# Assuming your main package is in the cmd/appname directory
go build -o nemo

echo "Build complete. Binary is in the 'build' directory."
