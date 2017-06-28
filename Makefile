ROOTFS_IMAGE		:= cirocosta/nfsvol-rootfs
ROOTFS_CONTAINER	:= rootfs
PLUGIN_NAME			:= nfsvol
PLUGIN_ID			:= $(shell docker plugin inspect $(PLUGIN_NAME) --format '{{ .ID }}')


install:
	cd main && \
		go install -v
	cd manager && \
		go install -v
	cd nfsvolctl && \
		go install -v


test:
	cd ./manager && go test -v


deps:
	glide install


fmt:
	cd ./main && gofmt -s -w .
	cd ./manager && gofmt -s -w .
	cd ./nfsvolctl && gofmt -s -w .


rootfs-image:
	docker build -t $(ROOTFS_IMAGE) .


rootfs: rootfs-image
	docker rm -vf $(ROOTFS_CONTAINER) || true
	docker create --name $(ROOTFS_CONTAINER) $(ROOTFS_IMAGE) true
	mkdir -p plugin/rootfs
	rm -rf plugin/rootfs/*
	docker export $(ROOTFS_CONTAINER) | tar -x -C plugin/rootfs
	docker rm -vf $(ROOTFS_CONTAINER)


plugin: rootfs
	docker plugin disable $(PLUGIN_NAME) || true
	docker plugin rm $(PLUGIN_NAME) || true
	docker plugin create $(PLUGIN_NAME) ./plugin
	docker plugin enable $(PLUGIN_NAME)


plugin-logs:
	docker run \
		--rm \
		-it \
		--privileged \
		--pid=host \
		justincormack/nsenter1 \
		/bin/sh -c 'docker-runc exec $(PLUGIN_ID) cat /var/log/nfsvol/plugin.log'


plugin-exec:
	docker run \
		--rm \
		-it \
		--privileged \
		--pid=host \
		justincormack/nsenter1 \
		/bin/sh -c 'docker-runc exec -t $(PLUGIN_ID) sh'


.PHONY: install deps fmt rootfs-image rootfs plugin plugin-logs plugin-exec test
