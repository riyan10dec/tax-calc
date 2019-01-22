package context

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig(path string) *Config {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("toml")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error context file: %s \n", err)
	}

	return &Config{
		DBHost:     config.Get("database.host").(string),
		DBPort:     config.Get("database.port").(string),
		DBUser:     config.Get("database.user").(string),
		DBPassword: config.Get("database.password").(string),
		DBName:     config.Get("database.name").(string),
	}
}
