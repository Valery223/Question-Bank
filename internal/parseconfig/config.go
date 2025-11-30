// Package parseconfig  отвечает за загрузку и парсинг конфигурации приложения
package parseconfig

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// Config представляет структуру конфигурации приложения
type Config struct {
	Env     string   `yaml:"env" env-default:"local"` // environment: local, dev, prod
	Address string   `yaml:"address" env-default:":8080"`
	DB      DBConfig `yaml:"db"`
}

// DBConfig содержит параметры подключения к базе данных
type DBConfig struct {
	Host string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	User string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	// Не указываем пароль и имя в yaml или указываем пустым
	Password string `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
	Name     string `yaml:"name" env:"DB_NAME" env-default:"question_bank"`
}

// MustLoad загружает конфигурацию из файла и переменных окружения и файла .env.
//
// Если путь к файлу не указан, используется "config/local.yaml" по умолчанию.
// В случае ошибки загрузки или парсинга конфигурации происходит паника.
func MustLoad() *Config {
	var configPath string
	flag.StringVar(&configPath, "config_path", "", "Path to the configuration file")
	flag.Parse()

	if configPath == "" {
		log.Println("No config path provided, using default 'config/local.yaml'")
		configPath = "config/local.yaml"
	}

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: No .env file found: %v", err)
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("configuration file does not exist: %s", configPath)
	}

	config := &Config{}
	if err := cleanenv.ReadEnv(config); err != nil {
		log.Printf("Warning: failed to read .env file: %v", err)
	}
	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	return config
}
