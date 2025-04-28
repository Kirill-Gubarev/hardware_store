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
	Host string `yaml:"host"`
	Port string `yaml:"port"`
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

	//read password
	passLen := len(cfg.Database.Password)
	if passLen >= 1 && cfg.Database.Password[0] == '$' {
		if passLen >= 3 && cfg.Database.Password[1] == '{' {
			cfg.Database.Password = os.Getenv(cfg.Database.Password[2:passLen-1])
		} else {
			cfg.Database.Password = os.Getenv(cfg.Database.Password[1:])
		}
	}
	config = cfg
	return cfg, nil
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
