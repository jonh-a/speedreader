OS_VERSIONS=( "linux" "windows" "darwin" )
ARCH_VERSIONS=( "arm64" "amd64" )

for os in "${OS_VERSIONS[@]}";
do
  for arch in "${ARCH_VERSIONS[@]}";
  do
    GOOS=$os GOARCH=$arch go build -o bin/speedread\_$os\_$arch;
  done;
done