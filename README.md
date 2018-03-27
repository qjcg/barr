The barr command prints out a status line for use with dwm(1).

# Dependencies

- `xsetroot` (arch package: `xort-xsetroot`)
- `iwgetid` for the wifi module (arch package: `wireless_tools`)


# Usage

```shell
# normal mode, uses `xsetroot` to set title for dwm.
barr

# test mode, prints a line to stdout for testing
barr -t
```
