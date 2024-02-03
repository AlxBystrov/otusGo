package main

import (
	"log"
	"os"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger LoggerConf
	// TODO
}

type LoggerConf struct {
	Level string `yaml:"level"`
	// TODO
}

func NewConfig() Config {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("error occurred while reading config: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	return Config{}
}

// TODO
