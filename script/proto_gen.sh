#!/bin/bash
cd "$(dirname "$0")" || exit
ENV="dev"
docker run --rm \
  -v "$(pwd)"/../resource/proto/${ENV}:/defs \
  -v "$(pwd)"/../src/gen:/out \
  namely/protoc-all \
  -d /defs \
  -l go \
  -o /out
echo "✅ finished. output: $(pwd)/../src/gen"