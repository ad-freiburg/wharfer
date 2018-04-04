Restricted Docker Wrapper
=========================

**WARNING THIS IS VERY EXPERIMENTAL WITH NO CLAIM OF ACTUAL SECURITY**

This is a wrapper around the docker command that only allows some basic
commands and flags. This can be used in addition to [No Trivial Root for
Docker](https://github.com/ad-freiburg/docker-no-trivial-root) to further
restrict access. Also if used with `setgid` and the `docker` group it can be
used to give this restricted docker access to users not in the docker group.

Building
--------
Make sure you have a Go environment [set up](https://golang.org/doc/install)
then do

    go get github.com/ad-freiburg/dockwrap

Installing
----------

    sudo cp $GOPATH/bin/dockwrap /usr/bin/
    sudo chown root:docker /usr/bin/dockwrap 
    sudo chmod g+s /usr/bin/dockwrap`

Also *make sure that the executable is only writable by root*
