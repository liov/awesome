CGO_ENABLED=1 \
GOOS=darwin \
GOARCH=amd64 \
CGO_LDFLAGS="-mmacosx-version-min=${MACOS_MIN_VER} --sysroot ${MACOS_SDK_PATH} -F/System/Library/Frameworks -L/usr/lib" \
CC="zig cc -mmacosx-version-min=${MACOS_MIN_VER} -target x86_64-macos-gnu -isysroot ${MACOS_SDK_PATH} -iwithsysroot /usr/include -iframeworkwithsysroot /System/Library/Frameworks" \
CXX="zig c++ -mmacosx-version-min=${MACOS_MIN_VER} -target x86_64-macos-gnu -isysroot ${MACOS_SDK_PATH} -iwithsysroot /usr/include -iframeworkwithsysroot /System/Library/Frameworks" \
go build -trimpath -buildmode=pie -o dist/darwin-amd64  ./cmd/...


CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="zig cc -target x86_64-linux" CXX="zig c++ -target x86_64-linux" go build 