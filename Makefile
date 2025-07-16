# windows: 'choco install make'
.PHONY: run dev test db-main redis up

# Executa a aplicação Go
run:
	go run .

# Inicia o servidor com hot-reload (usando Air)
dev:
	air

# Executa todos os testes do projeto
test:
	go test ./...

# Sobe todos os serviços essenciais (DB + Redis)
up: db-main redis

# Sobe o serviço de banco de dados principal via Docker Compose
db-main:
	docker compose -f 'docker-compose.yml' up -d --build 'database'

# Sobe o serviço Redis (cache) via Docker Compose
redis:
	docker compose -f 'docker-compose.yml' up -d --build 'redis'