Wharfer
=======
**WARNING THIS IS VERY EXPERIMENTAL WITH NO CLAIM OF ACTUAL SECURITY**

`wharfer` (pronounced /wɔɹfɚ/ from wharf ≈ pier ≈ dock) is a wrapper around the
`docker` command that only allows some basic commands and flags with the goal
of enabling `docker` usage by students on shared Linux machines. In the future
we may add access control for removing and killing containers using Unix
accounts.  Wharfer should be used together with the ["No Trivial Root for
Docker"](https://github.com/ad-freiburg/docker-no-trivial-root) authorization
plugin though technically it works without it. Also if used with `setgid` and
the `docker` group it allows a restricted access to `docker` while allowing
full `docker` access for everyone in the `docker` group.

Build/Download
--------------
Make sure you have a Go environment [set up](https://golang.org/doc/install)
then do

    go get github.com/ad-freiburg/wharfer

Alternatively you can download binary releases
[here](https://github.com/ad-freiburg/wharfer/releases)

Setup
-----

    # For a build from source
    sudo cp $GOPATH/bin/wharfer /usr/bin/
    # or for the binary release
    wget https://github.com/ad-freiburg/wharfer/releases/download/v0.1.0/wharfer_$(uname -m).tar.bz2
    tar -xavf wharfer_$(uname -m).tar.bz2
    sudo cp wharfer_$(uname -m)/wharfer /usr/bin/wharfer

    sudo chown root:docker /usr/bin/wharfer
    sudo chmod g+s /usr/bin/wharfer

Also *make sure that the executable is only writable by root*

Using wharfer
-------------
`wharfer` tries to be a drop-in replacement of `docker` for simple tasks. Though
there are some differences.

- Due to the use of the Go `flag` package not all options have long and short
  forms e.g. there's only `-p` and not `--publish`
- `-it` which in `docker` is a combination of the `-i` and `-t` options is only
  one option in `wharfer`

### Running Containers

A simple ephermal (through `--rm`) container running the `busybox` shell can be
executed as follows

    wharfer run --rm -it --name wharfer_busybox busybox:latest

### Building Containers
Using the busybox container from the previous section we can also build
a custom image just like with `docker`


    tee hello.sh <<EOF
     #!/bin/sh
     echo 'Hello, World!'
     EOF

    tee Dockerfile <<EOF
     FROM busybox:latest
     COPY hello.sh /app/
     CMD ["/bin/sh", "/app/hello.sh"]
     EOF

    wharfer build -t hellobusy .
    wharfer run --rm hellobusy

### Supported commands

The following `docker` commands are currently supported in some form

- `docker run` ⇒ `wharfer run`
- `docker build` ⇒ `wharfer build`
- `docker logs` ⇒ `wharfer logs`
- `docker ps` ⇒ `wharfer ps`
- `docker kill` ⇒ `wharfer kill`

To see the supported flags run `wharfer COMMAND --help`

A note on Volumes
-----------------
`wharfer run` supports the `-v` flag for mounting volumes (directories) inside
the container. However there are a few restrictions. Unlike with `docker` named
volumes are not supported and only mounting host directories through the `-v
/host/path:/container/path` syntax is allowed. As with `docker` only absolute
paths work.

When `wharfer` is used with user namespaces activated in the docker daemon (as
it should be)

**you need to make sure the permissions in your volumes are appropriately set**

For example in our default configuration `root` inside the container is mapped
to `nobody` outside the container. Thus if you want to write to a host
directory you need to make it writeable for `nobody`. Since a non-root user
can't change ownership of a directory the easiest way to make a local directory
writeable for `nobody` is to use `chown o+w hostdir.

An example using the busybox container goes as follows

    mkdir writetest
    chmod o+w writetest
    wharfer run --rm -it --name wharfer_busybox -v $(pwd)/writetest:/writetest busybox:latest
    # and then inside the container
    / # echo 'Hello, World!' > /writestest/hello.txt
    / # exit
    # and check the result on the host, the file hello.txt should be owned by
    # nobody
    ls -la writetest
    cat writetest/hello.txt
