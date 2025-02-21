package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true" `
}

// env-default:"production"
type Config struct {
	Env         string `yaml:"env" env: "ENV" env-required:"true" `
	StoragePath string `yam;: "storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func Mustload() *Config {
	var ConfigPath string

	ConfigPath = os.Getenv("CONFIG_PATH")
	if ConfigPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		ConfigPath = *flags

		if ConfigPath == "" {
			log.Fatal("config path is required")
		}
	}
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", ConfigPath)

	}
	var cfg Config
	err := cleanenv.ReadConfig(ConfigPath, &cfg)
	if err != nil {
		log.Fatalf("can't read config file: %s", err.Error())
	}
	return &cfg
}
