VERSION=0.1.0
PKG_NAME=tumblr

.PHONY: help
help: ## This help dialog.
	@IFS=$$'\n' ; \
    help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
    printf "%-30s %s\n" "target" "help" ; \
    printf "%-30s %s\n" "------" "----" ; \
    for help_line in $${help_lines[@]}; do \
        IFS=$$':' ; \
        help_split=($$help_line) ; \
        help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
        help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
        printf '\033[36m'; \
        printf "%-30s %s" $$help_command ; \
        printf '\033[0m'; \
        printf "%s\n" $$help_info; \
    done

.PHONY: tools
tools: ## Install testing tool packages
	GO111MODULE=off go get -u golang.org/x/tools/cmd/cover
	GO111MODULE=off go get -u golang.org/x/lint/golint	

.PHONY: fmt
fmt: ## Fmt fixer
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)

.PHONY: fmtcheck
fmtcheck: tools ##  Fmt check validation
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

.PHONY: lint
lint: fmtcheck ## Launch linter
	${HOME}/go/bin/golint -set_exit_status ./$(PKG_NAME)/

.PHONY: test
test: fmtcheck ## Launch tests
	go test -v ./$(PKG_NAME)/

.PHONY: testacc
testacc: fmtcheck ## Launch acc tests
	TF_ACC=1 go test -v ./$(PKG_NAME)/

.PHONY: cover
cover: fmtcheck ## Launch acc tests and calculate coverage
	TF_ACC=1 go test -v ./$(PKG_NAME)/ -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

.PHONY: codecov
codecov: fmtcheck ## Launch acc tests, calculate coverage and upload to codecov service.CODECOV_TOKEN env variable is needed.
	TF_ACC=1 go test -v ./$(PKG_NAME)/ -race -coverprofile=coverage.txt -covermode=atomic
	curl -s https://codecov.io/bash | bash

.PHONY: build
build: ## Build packages and dependencies
	rm -fr bin
	mkdir -p bin/windows_amd64
	mkdir -p bin/linux_amd64
	mkdir -p bin/darwin_amd64

	GOARCH=amd64 GOOS=windows go build -o bin/windows_amd64/terraform-provider-$(PKG_NAME)_v${VERSION}.exe
	GOARCH=amd64 GOOS=linux go build -o bin/linux_amd64/terraform-provider-$(PKG_NAME)_v${VERSION}
	GOARCH=amd64 GOOS=darwin go build -o bin/darwin_amd64/terraform-provider-$(PKG_NAME)_v${VERSION}

