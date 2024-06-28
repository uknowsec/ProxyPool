package Config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Interval int `yaml:"Interval"`
	ProxyApi struct {
		URL string `yaml:"url"`
	} `yaml:"ProxyApi"`
	Gost struct {
		Socks5User string `yaml:"socks5User"`
		Socks5Pass string `yaml:"socks5Pass"`
		Socks5Port int    `yaml:"socks5Port"`
		ApiURL     string `yaml:"apiurl"`
	} `yaml:"Gost"`
}

// LoadConfig loads YAML configuration from a file and returns a Config struct
func LoadConfig(filename string) (*Config, error) {
	// Read YAML file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse YAML data
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	log.Printf("Interval: %d", config.Interval)
	log.Printf("Proxy API URL: %s", config.ProxyApi.URL)
	log.Printf("SOCKS5 Username: %s", config.Gost.Socks5User)
	log.Printf("SOCKS5 Password: %s", config.Gost.Socks5Pass)
	log.Printf("SOCKS5 Port: %d", config.Gost.Socks5Port)
	log.Printf("API URL: %s", config.Gost.ApiURL)

	return &config, nil
}
