package config

import (
	"github.com/spf13/viper"
)

type api struct {
	Address         string `yaml:"address"`
	LogHTTPRequests bool   `yaml:"logHTTPRequests"`
}

type APIPConfigProvider interface {
	GetAPIAddress() string
	GetLogHttpRequests() bool
}

func (c *Config) GetAPIAddress() string {
	return c.API.Address
}

func (c *Config) GetLogHttpRequests() bool {
	return c.API.LogHTTPRequests
}

func (c *Config) OverrideAPIConfig() {
	if viper.IsSet("api-public-address") {
		c.API.Address = viper.GetString("api-public-address")
	}
}
