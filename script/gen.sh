#!/bin/sh

export MSYS_NO_PATHCONV=1
INPUT_FILE=${1:-openapi.yaml}
OUTPUT_DIR=${2:-gen}
PACKAGE_NAME=${3:-model}

rm -rf "${OUTPUT_DIR:?}"/*

docker run --rm -v "$(pwd):/local" openapitools/openapi-generator-cli generate \
    -i /local/"$INPUT_FILE" \
    -g go \
    -o /local/"$OUTPUT_DIR" \
    --global-property models \
    --additional-properties=packageName="$PACKAGE_NAME"

mkdir -p "$OUTPUT_DIR/$PACKAGE_NAME"

for file in gen/model_*.go; do
    new_name=$(basename "$file" | sed 's/^model_//')
    mv "$file" "$OUTPUT_DIR/$PACKAGE_NAME/$new_name"
done
