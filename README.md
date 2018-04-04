Wharfer
=======
**WARNING THIS IS VERY EXPERIMENTAL WITH NO CLAIM OF ACTUAL SECURITY**

`wharfer` (pronounced /wɔɹfɚ/ from wharf ≈ pier ≈ dock) is a wrapper around the
docker command that only allows some basic commands and flags with the goal of
enabling `docker` usage by students on shared Linux machines. In the future we
may add support for access control using the Unix user running the command.
Wharfer should be used together with the ["No Trivial Root for
Docker"](https://github.com/ad-freiburg/docker-no-trivial-root) authorization
plugin though technically it works without it. Also if used with `setgid`
and the `docker` group it allows a restricted access to docker while allowing
full docker access for everyone in the docker group.

Building
--------
Make sure you have a Go environment [set up](https://golang.org/doc/install)
then do

    go get github.com/ad-freiburg/wharfer

Installing
----------

    sudo cp $GOPATH/bin/wharfer /usr/bin/
    sudo chown root:docker /usr/bin/wharfer 
    sudo chmod g+s /usr/bin/wharfer`

Also *make sure that the executable is only writable by root*

Running
------
wharfer tries to be a drop-in replacement of docker for simple tasks. Though
there are some differences.

- Due to the use of the Go `flag` package not all options have long and short
  forms e.g. there's only `-p` and not `--publish`
- `-it` which in `docker` is a combination of the `-i` and `-t` options is only
  one option in `wharfer`

A simple ephermal (through `--rm`) container running the `busybox` shell can be
executed as follows

    wharfer run --rm -it --name whafer_busybox busybox:latest

The following `docker` commands are currently supported in some form

- `docker run` ⇒ `wharfer run`
- `docker build` ⇒ `wharfer build`
- `docker logs` ⇒ `wharfer logs`
- `docker ps` ⇒ `wharfer ps`
- `docker kill` ⇒ `wharfer kill`

To see the supported flags run `wharfer COMMAND --help`

