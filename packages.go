// +build package

//go:generate go build
//go:generate upx barr
//go:generate holo-build --force --format=pacman holo.toml
//go:generate rm -f barr
package main
