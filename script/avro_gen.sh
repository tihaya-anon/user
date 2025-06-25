#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")" || exit 1

cd ../resource || exit 1

PROTO_DIR="./proto"
AVRO_PARENT="./avro"

mkdir -p "${AVRO_PARENT}"

shopt -s nullglob                
for proto_path in "${PROTO_DIR}"/*.proto; do
  proto_file="$(basename "${proto_path}")"          
  proto_name="${proto_file%.proto}"                 
  out_dir="${AVRO_PARENT}/${proto_name}"            

  mkdir -p "${out_dir}"

  echo "▶ generate Avro schema: ${proto_file} → ${out_dir}"
  protoc \
    --proto_path="${PROTO_DIR}" \
    --avro_out="${out_dir}" \
    "${proto_path}"
done

echo "✅ finished. output: ${AVRO_PARENT}"
