.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/FRECNET/cmd/swarm
	go build -o bin/FRE     github.com/FRECNET/cmd/FRE
	go build -o bin/bootnode github.com/FRECNET/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/FRECNET/swarm/...
