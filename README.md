# nfsvol [![Build Status](https://travis-ci.com/cirocosta/nfsvol.svg?token=ixZ9XiEPW4YH62ixq7Av&branch=master)](https://travis-ci.com/cirocosta/nfsvol)

> Docker Volume Plugin designed to allow creating named volumes in a given path mounted as nfs

## Install

```
docker plugin install \
        --alias nfsvol \
        --grant-all-permissions \
        cirocosta/nfsvol
```


## Mounting Inside or Outside?

This plugin assumes that a NFS moint-point exists in the host. The rationale for keeping it outside the plugin container is that this way we can keep the NFS statistics going through `node_exporter`. 

In the TODO list we could extend this plugin to support mounting NFS at start up time.

