TARGET_DIR  := ./target
TARGET_NAME := nstu-cs
TARGET_PATH := ${TARGET_DIR}/${TARGET_NAME}

SRC_DIR := ./cmd/app

.PHONY: app
app: build run

.PHONY: run
run:
	${TARGET_PATH}

.PHONY: build
build: prepare
	go build -o ${TARGET_PATH} ${SRC_DIR}

.PHONY: prepare
prepare:
	go mod download
	@if [[ ! -d ${TARGET_DIR} ]]; then \
		mkdir ${TARGET_DIR}; \
	fi
