Restricting Wrapper Around docker
=================================

This is a wrapper around the docker command that only allows some basic
commands and flags. This can be used in addition to [No Trivial Root for
Docker](https://github.com/ad-freiburg/docker-no-trivial-root) to further
restrict access. Also if used with `setgid` and the `docker` group this can be
used to give very restricted docker access to users not in the docker group.

To build run

    go build

Then to install 

    sudo cp dockwrap /usr/bin/
    sudo chown root:docker /usr/bin/dockwrap 
    sudo chmod g+s /usr/bin/dockwrap`

Also *make sure that the executable is only writable by root*
