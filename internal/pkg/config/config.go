package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type ServerConfig struct {
	BindIP            string `yaml:"bind_ip"`
	Port              string `yaml:"port" env-default:"8080"`
	WriteTimeout      int    `yaml:"write_timeout"`
	ReadTimeout       int    `yaml:"read_timeout"`
	ReadHeaderTimeout int    `yaml:"read_header_timeout"`
	ShutdownTimeout   int    `yaml:"shutdown_timeout"`
}

type PostgresConfig struct {
	User   string `yaml:"user"`
	DBName string `yaml:"dbname"`
	// Password    string `env:"POSTGRES_PASSWORD"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	SSLmode     string `yaml:"sslmode"`
	MaxOpenCons int    `yaml:"max_open_cons"`
}

func NewConfig(configFile string) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	cfg := &Config{}
	err = cleanenv.ReadConfig(configFile, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
