BIN = barr

all:
	go build -o ${BIN} ./cmd/barr
	upx ${BIN}
	holo-build --format=pacman ./build/package/holo.toml

clean:
	rm -f ${BIN} $(wildcard *.pkg.tar.xz)
