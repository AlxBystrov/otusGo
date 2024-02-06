package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger  LoggerConf
	Storage StorageConf
	Server  ServerConf
}

type LoggerConf struct {
	Level   string `yaml:"level"`
	Handler string `yaml:"handler"`
}

type StorageConf struct {
	Type string `yaml:"type"`
}

type ServerConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewConfig() Config {
	config := &Config{}
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("error occurred while reading config: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		log.Fatalf("error occurred while decoding yaml: %s\n", err)
		os.Exit(1)
	}
	return *config
}
