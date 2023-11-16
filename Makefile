PROJECTNAME := $(shell basename "$(PWD)")
OS := $(shell uname -s | awk '{print tolower($$0)}')
GOARCH := amd64

## run: execete main application in local machine
.PHONY: run
run:
	go run cmd/main.go

## tidy: special go mod tidy without golang database checksum(GOSUMDB) 
.PHONY: tidy
tidy:
	export GOSUMDB=off ; go mod tidy

## test: run go test
test:
	go test -v ./...

## set_private_repo_global: set a "gitdev.devops.krungthai.com" to be a private repo in go global environment 
set_private_repo_global:
	go env -w GOPRIVATE="gitdev.devops.krungthai.com/*"

## update_standard_lib: update standard library (glo-standard-library) with GOPRIVATE option
update_standard_lib:
	GOPRIVATE=gitdev.devops.krungthai.com/glo/glo-standard-library go get gitdev.devops.krungthai.com/glo/glo-standard-library

## proto: generate proto files for gRPC
.PHONY: proto
proto:
	protoc internal/handlers/grpc/proto/*.proto --go_out=internal/handlers/grpc/pb --proto_path=internal/handlers/grpc/proto --go_opt=paths=source_relative --go-grpc_out=internal/handlers/grpc/pb --proto_path=internal/handlers/grpc/proto --go-grpc_opt=paths=source_relative

## up: docker compose up
.PHONY: up
up:
	docker-compose up -d

## down: docker compose down
.PHONY: down
down:
	docker-compose down

## init_folder_for_docker_compose: initial folder for docker-compose running
init_folder_for_docker_compose:
	mkdir -p {./data/kafka_data,./data/sql_data,./data/zookeeper}

## remove_data_for_kafka_error: remove all of data in kafka when docker-compose up error
remove_data_for_kafka_error:
	rm -rf ./data/kafka_data/* ./data/zookeeper/*

## help: helper
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Project: ["$(PROJECTNAME)"]"
	@echo " Please choose a command"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## gosec: run for scan code vulnerability by securego/gosec
.PHONY: gosec 
gosec: 
	gosec -exclude=G402 ./... 

## govulncheck: run for scan vulnerability package from Go vulnerability database
.PHONY: govulncheck
govulncheck: 
	govulncheck ./... 


## security: run make gosec and make govulncheck
security: gosec govulncheck