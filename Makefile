ROOTFS_IMAGE		:= cirocosta/nfsvol-rootfs
ROOTFS_CONTAINER	:= rootfs

install:
	cd main && \
		go install -v

deps:
	glide install

fmt:
	gofmt -s -w ./main

rootfs-image:
	docker build -t $(ROOTFS_IMAGE) .

rootfs: rootfs-image
	docker rm -vf $(ROOTFS_CONTAINER) || true
	docker create --name $(ROOTFS_CONTAINER) $(ROOTFS_IMAGE) true
	mkdir -p plugin/rootfs
	rm -rf plugin/rootfs/*
	docker export $(ROOTFS_CONTAINER) | tar -x -C plugin/rootfs
	docker rm -vf $(ROOTFS_CONTAINER)


.PHONY: install deps fmt rootfs-image rootfs
