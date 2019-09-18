VERSION=0.0.1
PKG_NAME=tumblr

default: build deploy

tools:
	GO111MODULE=off go get -u golang.org/x/tools/cmd/cover
	GO111MODULE=off go get -u golang.org/x/lint/golint	
	GO111MODULE=off go get -u github.com/mattn/goveralls	

build: fmtcheck
	GOARCH=amd64 GOOS=windows go build -o terraform-provider-$(PKG_NAME)_windows_amd64.exe
	GOARCH=amd64 GOOS=linux go build -o terraform-provider-$(PKG_NAME)_linux_amd64
	GOARCH=amd64 GOOS=darwin go build -o terraform-provider-$(PKG_NAME)_darwin_amd64

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)

fmtcheck: tools
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lint: fmtcheck
	${HOME}/go/bin/golint -set_exit_status ./$(PKG_NAME)/
test: fmtcheck
	go test -v ./$(PKG_NAME)/
testacc: fmtcheck
	TF_ACC=1 go test -v ./$(PKG_NAME)/

cover: fmtcheck
	TF_ACC=1 go test -v ./$(PKG_NAME)/ -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

coveralls: fmtcheck
	TF_ACC=1 go test -v ./$(PKG_NAME)/ -covermode=count -coverprofile=coverage.out
	${HOME}/go/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken ${COVERALLS_TOKEN}
	rm coverage.out

deploy: release
ifeq ($(OS),Windows_NT)
	cp bin/terraform-provider-$(PKG_NAME)_windows_amd64.exe $(dir $(shell which terraform))
else
	cp bin/terraform-provider-$(PKG_NAME)_darwin_amd64 $(dir $(shell which terraform))
endif


release: test
	rm -fr bin
	mkdir -p bin/windows_amd64
	mkdir -p bin/linux_amd64
	mkdir -p bin/darwin_amd64

	GOARCH=amd64 GOOS=windows go build -o bin/windows_amd64/terraform-provider-$(PKG_NAME)_v${VERSION}.exe
	GOARCH=amd64 GOOS=linux go build -o bin/linux_amd64/terraform-provider-$(PKG_NAME)_v${VERSION}
	GOARCH=amd64 GOOS=darwin go build -o bin/darwin_amd64/terraform-provider-$(PKG_NAME)_v${VERSION}

	mkdir -p releases/
	zip releases/terraform-provider-$(PKG_NAME)_windows_amd64_v${VERSION}.zip bin/windows_amd64/terraform-provider-$(PKG_NAME)_v${VERSION}.exe
	zip releases/terraform-provider-$(PKG_NAME)_linux_amd64_v${VERSION}.zip bin/linux_amd64/terraform-provider-$(PKG_NAME)_v${VERSION}
	zip releases/terraform-provider-$(PKG_NAME)_darwin_amd64_v${VERSION}.zip bin/darwin_amd64/terraform-provider-$(PKG_NAME)_v${VERSION}

.PHONY: tools build lint test testacc cover coveralls fmtcheck fmt deploy release build deps gets
