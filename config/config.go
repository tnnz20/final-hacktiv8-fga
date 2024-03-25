package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type PostgresConfig struct {
	Host     string `mapstructure:"Host"`
	Port     int    `mapstructure:"Port"`
	Name     string `mapstructure:"Name"`
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	SSLMode  string `mapstructure:"SSLMode"`
}
type ServerConfig struct {
	Name string `mapstructure:"Name"`
	Port int    `mapstructure:"Port"`
}

type TokenConfig struct {
	JWTSecret string `mapstructure:"JWTSecret"`
}
type DatabaseConfig struct {
	Postgres PostgresConfig
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Token    TokenConfig
}

func LoadConfig(env string) (*Config, error) {
	viper.SetConfigFile(fmt.Sprintf("config/config-%s.yaml", env))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
		return &Config{}, err
	}

	var config *Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshaling config: %v\n", err)
		return &Config{}, err
	}
	return config, nil
}

func (c *Config) GetDatabaseConfig() *DatabaseConfig {
	return &c.Database
}

func (c *Config) GetServerConfig() *ServerConfig {
	return &c.Server
}

func (c *Config) GetTokenConfig() *TokenConfig {
	return &c.Token
}
