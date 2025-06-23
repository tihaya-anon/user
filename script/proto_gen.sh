docker run --rm \
  -v $(pwd)/../resource/proto:/defs \
  -v $(pwd)/../src/gen:/out \
  namely/protoc-all \
  -f *.proto \
  -l go \
  -o /out
