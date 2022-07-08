.PHONY: test
test:
	@go test -mod=vendor -v -race ./...

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: vendor
vendor:
	@go mod vendor

.PHONY: build
build: fmt test
build:
	@mkdir -p dist/darwin-amd64
	@mkdir -p dist/linux-amd64
	@mkdir -p dist/windows-amd64
	@GOOS=darwin \
	GOARCH=amd64 \
	go build -o dist/darwin-amd64/passmanage ./cmd
	@GOOS=linux \
	GOARCH=amd64 \
	CGO_ENABLED=0 \
	go build -o dist/linux-amd64/passmanage ./cmd
	@GOOS=windows \
	GOARCH=amd64 \
	CGO_ENABLED=0 \
	go build -o dist/windows-amd64/passmanage ./cmd

# change the OS folder to test on your desired OS
.PHONY: create
create:
	./dist/darwin-amd64/passmanage create anon

.PHONY: get
get:
	./dist/darwin-amd64/passmanage get anon

.PHONY: list
list:
	./dist/darwin-amd64/passmanage list

.PHONY: generate
generate:
	./dist/darwin-amd64/passmanage generate