package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// [https://forum.golangbridge.org/t/getting-errors-when-running-go-project-that-uses-github-com-mattn-go-sqlite3-library/31800/2]
// baixar o compilador do gcc (C compiler) atraves do `choco`

func Conn_Sqlite() *sql.DB {
	var Conn, err = sql.Open("sqlite3", "./internal/db/database.db")
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}

	// Verificar se há resposta a conexão
	if err := Conn.Ping(); err != nil {
		fmt.Printf("Erro: %v", err)
		return nil
	}

	return Conn
}
