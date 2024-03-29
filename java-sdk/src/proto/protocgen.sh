proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  protoc \
  -I "." \
  --java_out=../main/java/. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

done