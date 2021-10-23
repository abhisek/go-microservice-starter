FROM golang:1.17-stretch AS builder

WORKDIR /tmp

RUN \
  apt-get update && \
  apt-get install -y unzip && \
  curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip && \
  unzip protoc-3.15.8-linux-x86_64.zip -d ${HOME}/.local && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1 && \
  go install github.com/envoyproxy/protoc-gen-validate@v0.6.2

WORKDIR /app

COPY . .

RUN PATH="${PATH}:${HOME}/.local/bin:${HOME}/go/bin" make
