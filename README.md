The barr command prints out a status line for use with dwm(1) and other
minimalistic window managers.

# Dependencies

- `xsetroot` (arch package: `xorg-xsetroot`)
- `iwgetid` for the wifi module (arch package: `wireless_tools`)


# Usage

```shell
# prints a status line to stdout
barr

# xsetroot mode, uses `xsetroot` to set title for dwm.
barr -x
```

# License

MIT.
