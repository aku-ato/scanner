include ../../common.mk

.PHONY: run, build, compile-proto

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
PROJECT_NAME := $(shell basename "$(ROOT_DIR)")
OUTPUT_DIR := $(PWD)/bin

COMMIT_SHA := $(shell git rev-parse --short HEAD 2>&1)
VER:='github.com/go-masonry/mortar/mortar.version=v1.2.3'
GIT:='github.com/go-masonry/mortar/mortar.gitCommit=${COMMIT_SHA}'
BUILD_TAG:='github.com/go-masonry/mortar/mortar.buildTag=42'
BUILD_TS:='github.com/go-masonry/mortar/mortar.buildTimestamp=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")'

run:
	@echo "Running $(PROJECT_NAME)"
	@go run -ldflags="-X ${VER} -X ${GIT} -X ${BUILD_TAG} -X ${BUILD_TS}" main.go config config/config.yml

build:
	@echo "Building $(PROJECT_NAME)"
	@go build -o $(OUTPUT_DIR)/$(PROJECT_NAME) -ldflags="-X ${VER} -X ${GIT} -X ${BUILD_TAG} -X ${BUILD_TS}" main.go

compile-proto:
	$(call compile_proto,$(PROJECT_NAME),$(ROOT_DIR))
	$(call add_to_git,"$(ROOT_DIR)/pkg/proto/","Update common proto files")
	$(call add_to_git,"$(ROOT_DIR)/third_party/OpenAPI/","Update common proto files")
	$(call add_to_git,"$(ROOT_DIR)/../../frontend/proto/scanner","Update common proto files")
