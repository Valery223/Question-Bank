# --- Этап 1: Сборка (Builder) ---
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

# 1. Сначала копируем только файлы зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# 2. Теперь копируем весь исходный код проекта
COPY . .

# 3. Собираем бинарный файл
# CGO_ENABLED=0 - выключает зависимость от системных C-библиотек (важно для Alpine)
# GOOS=linux - собираем под Linux
# -o /bin/server - куда положить результат
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/server cmd/server/main.go


# --- Этап 2: Запуск (Runner) ---
# Берем минимальный Linux (Alpine)
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 1. Копируем скомпилированный бинарник из первого этапа
COPY --from=builder /bin/server .

# 2. Копируем конфиг
COPY --from=builder /app/config/local.yaml ./config/local.yaml

# Миграции копировать не нужно, они уже зашиты внутри бинарника через embed

EXPOSE 8080

CMD ["./server"]