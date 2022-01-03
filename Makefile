.PHONY: build_mac build_win build_linux

PRO_NAME=go-cli
export CGO_ENABLED=0
export GOARCH=amd64
#windows,linux,darwin
export GOOS=windows

build:
ifeq ($(GOOS),windows)
	@go build -o $(PRO_NAME).exe
else
	@go build -o $(PRO_NAME)
endif

mod:
	@go mod tidy
