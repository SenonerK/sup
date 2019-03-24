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

case "$1" in
    "build")
        cd "$2"
        GOOS=linux go build -o $(basename "$2").bin
    ;;

    "proto")
        cd "$2"
        protoc --go_out=. --micro_out=. *.proto
    ;;

    "docker")
        docker-compose up -d --build $(basename "$2")_$(dirname "$2")
    ;;

    *)
        usage
    ;;
esac