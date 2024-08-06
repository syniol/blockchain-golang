build:
	go run ./cmd/ds_key_generator.go

test:
	go test -v ./...

up:
	docker-compose -f deploy/docker-compose.yml up -d

down:
	docker-compose -f deploy/docker-compose.yml down

shell:
	docker-compose -f deploy/docker-compose.yml exec app sh
