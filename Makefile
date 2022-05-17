.PHONY: all

#export VERSION := $(shell cat VERSION)
export BUILD_TIME := $(shell date --utc --rfc-3339 ns 2> /dev/null | sed -e 's/ /T/')
export GIT_COMMIT := $(shell git rev-parse --short HEAD 2> /dev/null || true)

export BIN_PATH := $(shell pwd)/bin

all:
	go build \
		-o ${BIN_PATH}/takler_client \
		main.go