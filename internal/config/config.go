package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Api      ApiConfig      `mapstructure:"api"`
	JWT      JwtConfig      `mapstructure:"jwt"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type ApiConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JwtConfig struct {
	EcdsaPrivateKey string `mapstructure:"ecdsa_private_key"`
	EcdsaPublicKey  string `mapstructure:"ecdsa_public_key"`
}

func LoadConfig(configPath string) (*Config, error) {
	// Set up Viper
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Unmarshal config into struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
