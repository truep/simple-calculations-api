app_name = calc-api
src_path = ./cmd/api/*.go
bin_path = ./bin/
build_flags = -ldflags "-s -w" 

build:
	go get -u ./...
	go mod tidy 
	go mod vendor
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(build_flags) -o $(bin_path)$(app_name) $(src_path)

.PHONY: clean

test: 
	go test ./...

clean:
	rm -rf $(bin_path)

all: test build