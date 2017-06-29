<h1 align="center">nfsvol 📂  </h1>

<h5 align="center">Docker Volume Plugin for managing NFS storage</h5>

<br/>


## Quickstart

The plugin assumes there's a directory in the host named `/mnt/nfs` which is where you have NFS mounted to.
This doesn't mean that you *really* need to have that as a NFS mount point. The plugin will put files there; not having NFS will only mean that it won't be distributed across hosts.

(are you a `docker-for-mac` user? Issue `docker run --rm --privileged --pid=host justincormack/nsenter1 /bin/sh -c 'mkdir -p /mnt/nfs'` to create `/mnt/nfs` inside the docker-for-mac VM. For more about the command, see #Tips section).


Having that, install the plugin:

```sh
docker plugin install \
        --grant-all-permissions \
        --alias nfsvol \
        cirocosta/nfsvol
```

If nothing went bad (Docker would complain) you can create named volumes and use them:

```sh
docker volume create \
        --driver nfsvol \
        myvolume1

docker run \
        -it \
        -v myvolume2:/somewhere \
        alpine \
        echo 'heyhey' > /somewhere/file2.txt

docker run \
        -it \
        -v myvolume3:/somewhere \
        alpine \
        echo 'heyhey' > /somewhere/file1.txt
```


The effect of this is having the data under a well defined structure under your NFS mount:


```
/mnt
  /nfs
    /myvolume1
    /myvolume2
      /file2.txt
    /myvolume3
      /file3.txt
```

## Tips (Docker-for-{mac,aws})

You can check whether you have a NFS mount active by using the `mount` command on Linux. For instance, using AWS EFS as the NFS mount point:

```sh
mount

/dev/xvdb1 on / type ext4 (rw,relatime,data=ordered)
/dev/xvdb1 on /mnt type ext4 (rw,relatime,data=ordered)
proc on /proc type proc (rw,nosuid,nodev,noexec,relatime)
tmpfs on /dev type tmpfs (rw,nosuid,mode=755)

...

# what we want:
<fs-id>.efs.us-west-2.amazonaws.com:/ on /mnt/nfs type nfs4 (rw,relatime,vers=4.1,rsize=1048576,wsize=1048576,namlen=255,hard,proto=tcp,timeo=600,retrans=2,sec=sys,clientaddr=<address>,local_lock=none,addr=<address>)
```

In case you're running `docker-for-mac` or `docker-for-aws` you don't have direct access to the host out of the box. For both cases you can use `justincormack/nsenter1` alongside `--pid` argument to enter the namespaces used for the pid 1 of the machine, giving you access to the filesystem:

```sh
docker run \
        --rm \
        -it \
        --privileged \
        --pid=host \
        justincormack/nsenter1 \
        /bin/sh
```

Having done that, just `mkdir -p /mnt/nfs`.


## Mounting Inside or Outside?

This plugin assumes that a NFS moint-point exists in the host. The rationale for keeping it outside the plugin container is that this way we can keep the NFS statistics going through `node_exporter`. 

In the TODO list we could extend this plugin to support mounting NFS at start up time and make more sense to have `nfs` in the name 🙌


[![Build Status](https://travis-ci.com/cirocosta/nfsvol.svg?token=ixZ9XiEPW4YH62ixq7Av&branch=master)](https://travis-ci.com/cirocosta/nfsvol)

