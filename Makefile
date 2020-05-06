include .env
export $(shell sed 's/=.*//' .env)

TARGET := pm2nb
SRC := cmd/main.go

.PHONY: build  generate run ${TARGET}

all: build

build: ${TARGET}

generate: swagger.json
	swagger generate client -f $^ -c internal/netbox

run: build
	./${TARGET}

${TARGET}: ${SRC}
	go build -o $@ $^

swagger.json:
	curl https://netbox.proelbtn.com/api/docs/?format=openapi > $@