package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	Env      string         `yaml:"env" env-default:"dev"`
	Server   ServerConfig   `yaml:"server" env-required:"true"`
	Database DatabaseConfig `yaml:"database" env-required:"true"`
}
type ServerConfig struct {
	GRPC GRPCConfig `yaml:"grpc"`
	HTTP HTTPConfig `yaml:"http"`
}

type GRPCConfig struct {
	Host string        `yaml:"host" env-default:"0.0.0.0"`
	Port int           `yaml:"port" env-default:"50051"`
	TTL  time.Duration `yaml:"ttl" env-default:"10h"`
}

type HTTPConfig struct {
	Host string `yaml:"host" env-default:"0.0.0.0"`
	Port int    `yaml:"port" env-default:"8080"`
}

type DatabaseConfig struct {
	Dsn string `yaml:"dsn" env-required:"true"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	return MustLoadByPath(path)
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "config path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
		slog.Info("CONFIG_PATH", res)
	}

	return res
}

func MustLoadByPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found")
	}
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
