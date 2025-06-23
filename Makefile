# windows: 'choco install make' 
.PHONY: run dev test db-main redis

run:
	go run . 

dev:
	air

test:
	go test ./...

# rodar banco principal no docker
db-main:
	docker compose -f 'docker-compose.yml' up -d --build 'database'

redis:
	docker compose -f 'docker-compose.yml' up -d --build 'redis'