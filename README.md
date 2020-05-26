# barr

[![Go Report Card](https://goreportcard.com/badge/github.com/qjcg/barr)](https://goreportcard.com/report/github.com/qjcg/barr)

The barr command prints out a status line for use with minimalistic window managers.

## Features

- simple
- lightweight
- easy to extend (blocks simply need to implement the [fmt.Stringer](https://golang.org/pkg/fmt/#Stringer) interface)


## Install

Download the [latest release](https://github.com/qjcg/barr/releases/latest) or install using use `go get`:

```
go get github.com/qjcg/barr/cmd/barr
```


## Usage

### With i3 or Sway

Add a `status_command` line like this to your [i3](https://i3wm.org/) or [Sway](https://swaywm.org) config file:

```
bar {
	status_command barr
	...
}
```


## External Tool Dependencies

The `barr` command itself does not depend on any external tools, but some individual blocks do.

Specifically, the blocks below have external tool dependencies:

- volume: depends on `pactl`
- wifi: depends on `iw`


## License

[MIT](./LICENSE)
