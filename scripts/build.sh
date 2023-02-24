#!/usr/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# export CGO_FLAGS="-O -D__BLST_PORTABLE__"

if ! [[ "$0" =~ scripts/build.sh ]]; then
    echo "must be run from repository root"
    exit 255
fi

# Set default binary directory location
name="o1f99fVrUeosvzAY53EtXybV9dheAP4trjCBGKt521qkEM2wb"

mkdir -p ./build

echo "Building hypertimestamp in ./build/$name"
go build -o ./build/$name ./cmd/hypertimestamp

