# ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
ROOT_DIR := /srv/

build:
	cd lang && cargo build --release
	cp lang/target/release/liblang.so lang/
	echo 'ROOT_DIR is $(ROOT_DIR)'
	go build -ldflags="-r $(ROOT_DIR)lang" main.go
