BUILD_OUTPUT:=go-microservice-starter

all: clean build

build: proto app

proto:
	protoc -I ./contracts/services/pet_service/ \
		-I ./contracts/lib/ -I ./contracts/lib/protoc-gen-validate \
		--go_out=./gen/ \
		--go-grpc_out=./gen/ \
		--validate_out="lang=go:./gen/" \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--validate_opt=paths=source_relative \
		./contracts/services/pet_service/pet_service.proto

app:
	go build -o ${BUILD_OUTPUT}

clean:
	-rm -rf ${BUILD_OUTPUT}
	-rm -rf ./gen
	-mkdir -p ./gen
