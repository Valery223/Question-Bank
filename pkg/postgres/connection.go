package postgres

import (
	"database/sql"
	"fmt"

	// Импортируем драйвер (pgx для stdlib)
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewPostgresDB создает и возвращает подключение к базе данных Postgres
func NewPostgresDB(host, port, user, password, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// "pgx" - это название драйвера
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	// Проверяем, что соединение работает
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	// Настройки пула (оптимизация)
	// db.SetMaxOpenConns(25)
	// db.SetMaxIdleConns(25)
	// db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
