FROM golang:1.17-stretch

WORKDIR /tmp

RUN \
	go install github.com/fullstorydev/grpcui/cmd/grpcui@latest && \
	go install github.com/jnewmano/grpc-json-proxy@latest

ENV PATH "${PATH}:${HOME}/go/bin"

WORKDIR /work


