
.PHONY: run
run:
	go run cmd/main.go

## tidy: special go mod tidy without golang database checksum(GOSUMDB) 
.PHONY: tidy
tidy:
	 go mod tidy

## up: docker compose up
.PHONY: up
up:
	docker-compose up -d

## down: docker compose down
.PHONY: down
down:
	docker-compose down
