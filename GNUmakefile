VERSION=0.0.1
PKG_NAME=tumblr

default: build deploy

gets:
	go get github.com/hashicorp/terraform
	go get github.com/stretchr/testify
	go get gopkg.in/alecthomas/gometalinter.v2

deps:
	go install github.com/hashicorp/terraform
	go install github.com/stretchr/testify

build: fmtcheck
	GOARCH=amd64 GOOS=windows go build -o terraform-provider-tumblr_windows_amd64.exe
	GOARCH=amd64 GOOS=linux go build -o terraform-provider-tumblr_linux_amd64
	GOARCH=amd64 GOOS=darwin go build -o terraform-provider-tumblr_darwin_amd64

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

test: fmtcheck
	go test -v ./tumblr/
testacc: fmtcheck
	TF_ACC=1 go test -v ./tumblr/

cover: fmtcheck
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	TF_ACC=1 go test -v ./tumblr/ -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

deploy: release
ifeq ($(OS),Windows_NT)
	cp bin/terraform-provider-tumblr_windows_amd64.exe $(dir $(shell which terraform))
else
	cp bin/terraform-provider-tumblr_darwin_amd64 $(dir $(shell which terraform))
endif


release: test
	rm -fr bin
	mkdir -p bin/windows_amd64
	mkdir -p bin/linux_amd64
	mkdir -p bin/darwin_amd64

	GOARCH=amd64 GOOS=windows go build -o bin/windows_amd64/terraform-provider-tumblr_v${VERSION}.exe
	GOARCH=amd64 GOOS=linux go build -o bin/linux_amd64/terraform-provider-tumblr_v${VERSION}
	GOARCH=amd64 GOOS=darwin go build -o bin/darwin_amd64/terraform-provider-tumblr_v${VERSION}

	mkdir -p releases/
	zip releases/terraform-provider-tumblr_windows_amd64_v${VERSION}.zip bin/windows_amd64/terraform-provider-tumblr_v${VERSION}.exe
	zip releases/terraform-provider-tumblr_linux_amd64_v${VERSION}.zip bin/linux_amd64/terraform-provider-tumblr_v${VERSION}
	zip releases/terraform-provider-tumblr_darwin_amd64_v${VERSION}.zip bin/darwin_amd64/terraform-provider-tumblr_v${VERSION}

.PHONY: build test testacc cover fmtcheck fmt deploy release build deps gets