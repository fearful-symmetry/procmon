# Procmon

`procmon` is a CLI tool to view linux process events (fork, exit, etc) in real-time

This is  WIP. There will be more here soon.


## Building and running

`procmon` uses [modules](https://github.com/golang/go/wiki/Modules) for dependency management. This means you need at least go 1.11.

```bash

$ git clone github.com/fearful-symmetry/procmon
$ cd procmon
$ go build
# the proc connector netlink socket requires root
$ sudo ./procmon

```


## TODO

- Add a little signal number (17 == SIGCHLD) helper
- JSON output