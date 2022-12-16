ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

build:
    cd lang && cargo build --release
    cp lang/target/release/librustaceanize.dylib lib/
    echo 'ROOT_DIR is $(ROOT_DIR)'
    go build -ldflags="-r $(ROOT_DIR)lib" main.go