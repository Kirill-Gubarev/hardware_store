package utils

import (
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Service  ServiceConfig  `yaml:"service"`
	Database DatabaseConfig `yaml:"database"`
}

type ServiceConfig struct {
	Host    string  `yaml:"host"`
	Port    string  `yaml:"port"`
	APIKey  string  `yaml:"apiKey"`
}

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

var config *Config = nil

func GetConfig() (*Config, error) {
	if config != nil{
		return config, nil
	}

	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	var decoder = yaml.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		config = nil
		return nil, err
	}

	subStitution(&cfg.Database.Password)
	subStitution(&cfg.Service.APIKey)

	config = cfg
	return cfg, nil
}
func subStitution(field *string) {
	fieldLen := len(*field)
	if fieldLen >= 1 && (*field)[0] == '$' {
		if fieldLen >= 3 && (*field)[1] == '{' {
			*field = os.Getenv((*field)[2 : fieldLen-1])
		} else {
			*field = os.Getenv((*field)[1:])
		}
	}
}
func GetServiceConfig() (*ServiceConfig, error){
	cfg, err := GetConfig()
	if err != nil{
		return nil, err
	}
	return &cfg.Service, nil
}
func GetDatabaseConfig() (*DatabaseConfig, error){
	cfg, err := GetConfig()
	if err != nil{
		return nil, err
	}
	return &cfg.Database, nil
}
