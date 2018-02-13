GOFLAGS := --ldflags '-w -linkmode external'
CC := $(shell which musl-clang)

all: images

images: coverservice/coverservice-linux-amd64 \
		docker-compose.yml
	docker-compose build

COMMON_FILES := $(wildcard common/*.go)

coverservice/coverservice-linux-amd64: \
		$(COMMON_FILES) \
		coverservice/main.go \
		$(wildcard coverservice/service/*.go)
	cd coverservice && \
	CC=${CC} go build ${GOFLAGS} -o coverservice-linux-amd64
