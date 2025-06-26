#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")" || exit 1
cd ../resource || exit 1

ENV="prod"
PROTO_DIR="./proto/${ENV}"
AVRO_PARENT="./avro/${ENV}"
MAPPING_FILE="${AVRO_PARENT}/schema_registry_mapping.yaml"

mkdir -p "${AVRO_PARENT}"
rm -f "${MAPPING_FILE}"
touch "${MAPPING_FILE}"

{
  echo "# Auto-generated schema registry subject mapping"
  echo "# ENV: ${ENV}"
  echo "schemas:"
} >>"${MAPPING_FILE}"

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

  entries=""
  for avsc_file in "${out_dir}"/*.avsc; do
    msg_name="$(basename "${avsc_file}" .avsc)"
    service="${proto_name}"
    subject="${ENV}.${service}.${proto_name}.${msg_name}-value"

    entries="${entries}
  - proto: ${proto_file}
    message: ${msg_name}
    avsc_path: ${avsc_file#./}
    subject: ${subject}"
  done

  # 一次性写入本 proto 下的所有 message 映射，避免多次重定向
  {
    echo "  # ====================="
    echo "  # ${proto_name}"
    echo "  # =====================${entries}"
  } >>"${MAPPING_FILE}"
done

echo "✅ finished. output: ${AVRO_PARENT}"
echo "📄 mapping file: ${MAPPING_FILE}"
