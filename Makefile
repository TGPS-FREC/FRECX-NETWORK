.PHONY: FRE FRE-cross evm all test clean
.PHONY: FRE-linux FRE-linux-386 FRE-linux-amd64 FRE-linux-mips64 FRE-linux-mips64le
.PHONY: FRE-darwin FRE-darwin-386 FRE-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= 1.13.1
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

FRE:
	go run build/ci.go install ./cmd/FRE
	@echo "Done building."
	@echo "Run \"$(GOBIN)/FRE\" to launch FRE."

gc:
	go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	go run build/ci.go install

test: all
	go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

FRE-cross: FRE-windows-amd64 FRE-darwin-amd64 FRE-linux
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/FRE-*

FRE-linux: FRE-linux-386 FRE-linux-amd64 FRE-linux-mips64 FRE-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/FRE-linux-*

FRE-linux-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/FRE
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/FRE-linux-* | grep 386

FRE-linux-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/FRE
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/FRE-linux-* | grep amd64

FRE-linux-mips:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/FRE
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/FRE-linux-* | grep mips

FRE-linux-mipsle:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/FRE
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/FRE-linux-* | grep mipsle

FRE-linux-mips64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/FRE
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/FRE-linux-* | grep mips64

FRE-linux-mips64le:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/FRE
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/FRE-linux-* | grep mips64le

FRE-darwin: FRE-darwin-386 FRE-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/FRE-darwin-*

FRE-darwin-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/FRE
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/FRE-darwin-* | grep 386

FRE-darwin-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/FRE
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/FRE-darwin-* | grep amd64

FRE-windows-amd64:
	go run build/ci.go xgo -- --go=$(GO) -buildmode=mode -x --targets=windows/amd64 -v ./cmd/FRE
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/FRE-windows-* | grep amd64
gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
