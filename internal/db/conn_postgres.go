package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Conn_Postgres() *sql.DB {
	conn, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("Erro de conexão com banco: %v", err)
		return nil
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Erro de resposta: %v", err)
		return nil
	}

	return conn
}
