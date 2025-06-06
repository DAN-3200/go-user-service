package repository

import "fmt"

func (it *SQLManager) CreateUserTable() error {
	var _, err = it.DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			is_email_verified BOOLEAN DEFAULT FALSE,
			is_active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			role VARCHAR(50) NOT NULL
		);`,
	)
	if err != nil {
		fmt.Println("Erro:", err)
		return err
	}
	return nil
}
