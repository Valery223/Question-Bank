package postgres_test

import (
	"context"
	"database/sql"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// SetupTestDB поднимает настоящий Postgres в Докере и накатывает миграции
func SetupTestDB(t *testing.T) *sql.DB {
	ctx := context.Background()

	dbContainer, err := postgres.Run(ctx,
		"postgres:15-alpine",
		postgres.WithDatabase("test_db"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		// Ждем, пока порт БД станет доступен и пойдут логи
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	require.NoError(t, err)

	// Гарантируем, что контейнер умрет после теста (даже если тест упадет)
	t.Cleanup(func() {
		if err := dbContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	// 2. Получаем строку подключения
	connStr, err := dbContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	// 3. Подключаемся через драйвер
	db, err := sql.Open("pgx", connStr)
	require.NoError(t, err)

	// 4. Накатываем миграции (Goose)
	// Хитрость: нам нужно найти путь к папке migrations относительно этого файла.
	// Так как тест лежит в internal/repository/postgres, нам нужно подняться на 3 уровня вверх.

	// Получаем путь к текущему файлу
	_, filename, _, _ := runtime.Caller(0)
	// Строим путь: <project_root>/migrations
	migrationsDir := filepath.Join(filepath.Dir(filename), "../../app/migrations")

	// Накатываем
	require.NoError(t, goose.SetDialect("postgres"))
	err = goose.Up(db, migrationsDir)
	require.NoError(t, err, "failed to apply migrations")

	return db
}
