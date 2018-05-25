// +build package

//go:generate vgo build
//go:generate holo-build --force --format=pacman holo.toml
//go:generate rm -f barr
package main
