#!/bin/bash

function usage {
    echo "usage: $0 task path"
    echo -e "\ttask: build, proto"
    echo -e "\tpath: e.g. srv/auth"
    exit 1
}

# if wrong amount of arguments
if [ $# -lt 2 ]; then
    usage
fi

# if path arg is not a directory
if ! [ -d "$2" ]; then
    usage
fi

cd "$2"

case "$1" in
    "build")
        GOOS=linux go build -o $(basename "$2").bin
    ;;

    "proto")
        protoc --go_out=. --micro_out=. *.proto
    ;;

    *)
        usage
    ;;
esac