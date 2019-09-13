version=0.0.1

default: build deploy

gets:
	go get github.com/hashicorp/terraform
	go get github.com/stretchr/testify
	go get gopkg.in/alecthomas/gometalinter.v2

deps:
	go install github.com/hashicorp/terraform
	go install github.com/stretchr/testify

build:
	GOARCH=amd64 GOOS=windows go build -o terraform-provider-tumblr_windows_amd64.exe
	GOARCH=amd64 GOOS=linux go build -o terraform-provider-tumblr_linux_amd64
	GOARCH=amd64 GOOS=darwin go build -o terraform-provider-tumblr_darwin_amd64

test:
	go test -v ./tumblr/
testacc:
	TF_ACC=1 go test -v ./tumblr/

cover:
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

plan:
	@terraform plan

release: test
	rm -fr bin
	mkdir -p bin/windows_amd64
	mkdir -p bin/linux_amd64
	mkdir -p bin/darwin_amd64

	GOARCH=amd64 GOOS=windows go build -o bin/windows_amd64/terraform-provider-tumblr_v${version}.exe
	GOARCH=amd64 GOOS=linux go build -o bin/linux_amd64/terraform-provider-tumblr_v${version}
	GOARCH=amd64 GOOS=darwin go build -o bin/darwin_amd64/terraform-provider-tumblr_v${version}

	mkdir -p releases/
	zip releases/terraform-provider-tumblr_windows_amd64_v${version}.zip bin/windows_amd64/terraform-provider-tumblr_v${version}.exe
	zip releases/terraform-provider-tumblr_linux_amd64_v${version}.zip bin/linux_amd64/terraform-provider-tumblr_v${version}
	zip releases/terraform-provider-tumblr_darwin_amd64_v${version}.zip bin/darwin_amd64/terraform-provider-tumblr_v${version}
