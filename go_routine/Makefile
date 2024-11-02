.DEFAULT_GOAL := clean
.PHONE:fmt vet build clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

run: build
	./go_routine

clean: run
	go clean