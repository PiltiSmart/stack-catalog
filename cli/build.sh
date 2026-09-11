#!/usr/bin/env bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "${SCRIPT_DIR}"

echo "======================================================="
echo "   BUILDING THINGSBOARD 'ps' CLI (GO + COBRA)"
echo "======================================================="

mkdir -p bin
echo "[1/2] Downloading Go dependencies..."
go mod tidy

echo "[2/2] Compiling standalone static binary..."
go build -ldflags="-s -w" -o bin/ps main.go
chmod +x bin/ps

echo "[SUCCESS] Compiled binary available at: bin/ps"
echo ""
echo "To install globally, run:"
echo "   cp bin/ps /usr/local/bin/ps"
