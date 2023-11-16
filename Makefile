.PHONY: all

export BIN_PATH := $(shell pwd)/bin

all:
	go build \
		-o ${BIN_PATH}/takler_client \
		main.go