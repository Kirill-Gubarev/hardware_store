package utils

import (
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Service         ServiceConfig  `yaml:"service"`
	AddressService  ServiceConfig  `yaml:"addressService"`
	ProductService  ServiceConfig  `yaml:"productService"`
	UserService     ServiceConfig  `yaml:"userService"`
}

type ServiceConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (s ServiceConfig) EndPoint() string {
	return s.Host + ":" + s.Port
}
func (s ServiceConfig) URL() string {
	return "http://" + s.EndPoint()
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
func GetAddressServiceConfig() (*ServiceConfig, error){
	cfg, err := GetConfig()
	if err != nil{
		return nil, err
	}
	return &cfg.AddressService, nil
}
func GetProductServiceConfig() (*ServiceConfig, error){
	cfg, err := GetConfig()
	if err != nil{
		return nil, err
	}
	return &cfg.ProductService, nil
}
func GetUserServiceConfig() (*ServiceConfig, error){
	cfg, err := GetConfig()
	if err != nil{
		return nil, err
	}
	return &cfg.UserService, nil
}
