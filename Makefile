all: clean_create_dir pet_service

clean_create_dir:
	-rm -rf ./gen
	-mkdir -p ./gen

pet_service:
	protoc -I ./contracts/services/pet_service/ \
		-I ./contracts/lib/ -I ./contracts/lib/protoc-gen-validate \
		--go_out=./gen/ \
		--go-grpc_out=./gen/ \
		--validate_out="lang=go:./gen/" \
		./contracts/services/pet_service/pet_service.proto

clean:
	-rm -rf pet_service
	-rm -rf gen
