#!/bin/bash
cd "$(dirname "$0")" || exit

docker run --rm \
  -v "$(pwd)"/../resource/proto:/defs \
  -v "$(pwd)"/../src/gen:/out \
  namely/protoc-all \
  -d /defs \
  -l go \
  -o /out