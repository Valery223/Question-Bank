package parseconfig

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env     string   `yaml:"env" env-default:"local"` // environment: local, dev, prod
	Address string   `yaml:"address" env-default:":8080"`
	DB      DBConfig `yaml:"db"`
}

type DBConfig struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"password"`
	Name     string `yaml:"name" env-default:"question_bank"`
}

func MustLoad() *Config {
	var configPath string
	flag.StringVar(&configPath, "config_path", "", "Path to the configuration file")
	flag.Parse()

	if configPath == "" {
		log.Fatal("config_path flag is required")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("configuration file does not exist: %s", configPath)
	}

	config := &Config{}
	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	return config
}
