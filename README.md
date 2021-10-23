# Go Microservice Starter
An opinionated Microservice starter in Go

# Usage

Checkout the code.

# TLDR;

Start the service:

```bash
docker build -t go-microservice-start .
docker run --rm -it -p 8080:8080 go-microservice-start ./microservice_starter
```

Interact with the service using [grpcui](https://github.com/fullstorydev/grpcui):

```bash
./tools/run.sh grpcui -plaintext localhost:8080
```
