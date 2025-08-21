# Go User Service - Clean Architecture

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

## Descrição

API RESTful para gerenciamento de usuários e autenticação, desenvolvida em **Golang** utilizando o framework HTTP **Gin**, com banco de dados **PostgreSQL** e data cache **Redis**, seguindo o padrão **Clean Architecture**.

Este projeto teve como objetivo o aprendizado e a consolidação dos conceitos e ferramentas utilizadas.

## Tecnologias

| Tecnologia                                   | Descrição                |
| -------------------------------------------- | ------------------------ |
| [**Go**](https://golang.org)                 | Linguagem do back end    |
| [**Gin**](https://gin-gonic.com)             | Framework HTTP para Go   |
| [**PostgreSQL**](https://www.postgresql.org) | Banco de dados principal |
| [**Redis**](https://redis.io)                | Sessão de usuário        |
| [**Docker**](https://www.docker.com)         | Containers               |

## Estrutura do projeto

```bash
├── internal
│   ├── domain               # Regras de negócio e contratos da aplicação
│   │   ├── dto              # Data Transfer Objects
│   │   ├── entity           # Entidades do domínio
│   │   ├── ports            # Interfaces (contratos de entrada/saída)
│   │   └── usecase          # Casos de uso (regras de negócio aplicadas)
│   ├── infrastructure       # Implementação de interfaces externas
│   │   ├── adapters         # Adaptações para dependências externas (ex: repos, services)
│   │   ├── db               # Conexão ao banco
│   │   └── http             # Camada de entrega HTTP
│   │       ├── controller   # Controladores 
│   │       ├── middlewares  # Middlewares HTTP 
│   │       ├── routes       # Definição das rotas da API
│   │       └── server       # Inicialização do servidor HTTP
│   └── tests                # Testes internos
│       ├── integration      # Testes de integração 
│       └── unit             # Testes unitários 
├── pkg                      # Pacotes compartilhados entre módulos
│   ├── resttype             # Definições de tipos de resposta REST
│   └── utils                # Funções utilitárias genéricas
├── main.go                  # Ponto de entrada da aplicação
├── go.mod                   # Módulo e dependências do projeto
├── go.sum                   # Hashes das dependências
├── .air.toml                # Configuração do Air (live reload)
├── .dockerignore            # Arquivos ignorados pelo Docker
├── .env.example             # Exemplo de variáveis de ambiente
├── .gitignore               # Arquivos ignorados pelo Git
├── docker-compose.yml       # Orquestração de serviços com Docker
├── dockerfile               # Dockerfile da aplicação
├── LICENSE                  # Licença do projeto
├── Makefile                 # Atalhos para comandos e automações
├── prometheus.yml           # Configuração do Prometheus
└── README.md                # Documentação do projeto
```

## Instalação e execução

```bash
# Clonar repositório
git clone https://github.com/DAN-3200/go-user-service.git
cd go-user-service
```
Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:
```bash
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=database
POSTGRES_URL="postgresql://user:password@localhost:5400/database?sslmode=disable"

SECRET_KEY="SECRET_KEY"
```

```Makefile
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
```

###### Subir os serviços Docker (DB + Redis)
```bash
make up
```

###### Executar a aplicação
```bash
# Desenvolvimento com hot-reload
make dev

# Ou execução padrão
make run
```

## Endpoints da API

| Route                                                  | Descrição                                |
| ------------------------------------------------------ | ---------------------------------------- |
| <kbd>POST /users</kbd>                                 | Criar novo usuário (admin)               |
| <kbd>GET /users</kbd>                                  | Listar usuários (admin)                  |
| <kbd>GET /users/:id</kbd>                              | Buscar usuário por ID (admin)            |
| <kbd>PATCH /users/:id</kbd>                            | Atualizar usuário por ID (admin)         |
| <kbd>DELETE /users/:id</kbd>                           | Remover usuário por ID (admin)           |
| <kbd>POST /auth/login</kbd>                            | Login do usuário                         |
| <kbd>POST /auth/logout</kbd>                           | Logout do usuário                        |
| <kbd>POST /auth/register</kbd>                         | Registro de novo usuário                 |
| <kbd>POST /auth/refresh-token</kbd>                    | Gerar novo token de acesso               |
| <kbd>POST /auth/verify-email</kbd>                     | Verificar e-mail do usuário              |
| <kbd>GET /auth/forget-password/send-token/:email</kbd> | Enviar token de redefinição de senha     |
| <kbd>POST /auth/forget-password/refresh-password</kbd> | Redefinir senha com token                |
| <kbd>GET /me</kbd>                                     | Obter informações do usuário autenticado |
| <kbd>PATCH /me</kbd>                                   | Atualizar informações do próprio usuário |

## Roadmap

| Objetivo                        | Status                  | Previsão  |
| ------------------------------- | ----------------------- | --------- |
| Envio de token por e‑mail       | <kbd>concluido</kbd>    | Maio/2025 |
| Integração Prometheus & Grafana | <kbd>em andamento</kbd> | Ago/2025  |
| Pipeline de deploy (CI/CD)      | <kbd>pendente</kbd>     | Set/2025  |

## Licença

Este projeto está licenciado sob a Licença MIT. Consulte o arquivo [LICENSE](./LICENSE) para mais detalhes.
