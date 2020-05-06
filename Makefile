include .env
export $(shell sed 's/=.*//' .env)

TARGET := pm2nb
SRC := cmd/main.go

.PHONY: ${TARGET}

all: build

build: ${TARGET}

run: build
	./${TARGET}

${TARGET}: ${SRC}
	go build -o $@ $^
