# --- Переменные ---
# Название бинарника
APP_NAME=server
# Путь к главному файлу
MAIN_FILE=cmd/server/main.go
# Путь к конфигу
CONFIG_FILE=config/local.yaml

all: build

# --- 1. Установка зависимостей ---
# Устанавливаем тулзы (goose, mockery) в $GOPATH/bin
deps:
	@echo "Installing tools..."
	go install github.com/vektra/mockery/v3@v3.6.1
	go install github.com/pressly/goose/v3/cmd/goose@v3.26.0
	@echo "Tools installed."

# --- 2. Разработка и Запуск ---
# Сборка бинарника
build:
	@echo "Building..."
	go build -o bin/$(APP_NAME) $(MAIN_FILE)

# Просто запуск (go run)
run:
	go run $(MAIN_FILE) -config_path $(CONFIG_FILE)

# --- 3. Тестирование ---
# Запуск юнит-тестов (все, кроме интеграционных)
test:
	go test $(TEST_FLAGS) ./internal/... ./pkg/...

test-cover:
	go test -coverprofile=coverage.out ./internal/...
	go tool cover -html=coverage.out
	rm coverage.out

# Накатить миграции (вверх)
# migrate-up:
# 	goose -dir migrations postgres $(DB_DSN) up

# # Откатить последнюю миграцию (вниз)
# migrate-down:
# 	goose -dir migrations postgres $(DB_DSN) down

# Статус миграций
# migrate-status:
# 	goose -dir migrations postgres $(DB_DSN) status

# Создать новую миграцию. Использование: make migrate-create NAME=add_users
migrate-create:
	@if [ -z "$(NAME)" ]; then echo "ERR: NAME is not set. Use: make migrate-create NAME=my_migration"; exit 1; fi
	goose -dir internal/app/migrations create $(NAME) sql


# Линтер (проверка кода на чистоту)
lint:
	golangci-lint run

clean:
	rm -rf ./bin/

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down