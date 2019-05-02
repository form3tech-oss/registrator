NAME=registrator
VERSION=$(shell cat VERSION)
DEV_RUN_OPTS ?= consul:

test:
	go test ./bridge
	go test ./consul
	go test ./consulkv
	go test ./etcd
	go test ./skydns2
	go test ./zookeeper

build: bin/registrator

bin/registrator:
	mkdir -p bin
	go build -o bin/registrator .

release:
	goreleaser

docs:
	boot2docker ssh "sync; sudo sh -c 'echo 3 > /proc/sys/vm/drop_caches'" || true
	docker run --rm -it -p 8000:8000 -v $(PWD):/work gliderlabs/pagebuilder mkdocs serve

.PHONY: test release docs
