#!/usr/bin/env bash
# ==============================================================================
# Cross-compilation script for 'ps' CLI across Linux, macOS and Windows
# ==============================================================================

set -e

OUTPUT_DIR="dist"
mkdir -p "${OUTPUT_DIR}"

echo "=========================================================="
echo "  Building 'ps' CLI for Multiple OS & Architectures"
echo "=========================================================="

go mod tidy

PLATFORMS=(
    "linux/amd64/ps-linux-amd64"
    "linux/arm64/ps-linux-arm64"
    "darwin/amd64/ps-darwin-amd64"
    "darwin/arm64/ps-darwin-arm64"
    "windows/amd64/ps-windows-amd64.exe"
)

for target in "${PLATFORMS[@]}"; do
    IFS="/" read -r GOOS GOARCH OUTNAME <<< "$target"
    echo "Compiling for ${GOOS}/${GOARCH} -> ${OUTPUT_DIR}/${OUTNAME}..."
    env GOOS="${GOOS}" GOARCH="${GOARCH}" CGO_ENABLED=0 go build -ldflags="-s -w" -o "${OUTPUT_DIR}/${OUTNAME}" main.go
    chmod +x "${OUTPUT_DIR}/${OUTNAME}" || true
done

echo ""
echo "✔ All binaries compiled successfully in '${OUTPUT_DIR}/':"
ls -lh "${OUTPUT_DIR}"
