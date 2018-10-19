// +build packages

//go:generate go build
//go:generate upx barr
//go:generate holo-build --force --format=pacman ../../build/packages/holo.toml
//go:generate rm -f barr
package main
