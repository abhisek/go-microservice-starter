#!/bin/bash

cd `dirname $0`

docker build -t go-microservice-starter-tools .
docker run --rm -it --net host go-microservice-starter-tools $@
