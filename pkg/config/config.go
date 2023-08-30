package config

import (
	"io/ioutil"

	"github.com/upstars-global/go-service-skeleton/pkg/argumentsresolver"
	"gopkg.in/yaml.v2"
)

type GeneralConfigProvider interface {
	GetEnvName() Env
	DebugEnabled() bool
}

type Config struct {
	EnvName Env    `yaml:"envName"`
	Debug   bool   `yaml:"debug"`
	Logger  logger `yaml:"logger"`

	API api `yaml:"api"`
	DB  db  `yaml:"db"`

	configFile string
}

type Env string

const (
	EnvLocal Env = "local"
	EnvDev   Env = "develop"
	EnvStage Env = "staging"
	EnvProd  Env = "production"
)

func (e Env) String() string {
	return string(e)
}

func New(arguments argumentsresolver.ArgumentsInterface) (c *Config, err error) {
	c = &Config{}
	c.configFile, err = arguments.GetString(argumentsresolver.ArgumentConfigName)

	if err != nil {
		panic(err)
	}

	if err = c.loadConfig(); err != nil {
		return
	}
	c.OverrideDBConfig()
	c.OverrideAPIConfig()

	verbose, err := arguments.GetBool(argumentsresolver.ArgumentConfigVerbose)
	if err != nil {
		panic(err)
	}
	if verbose {
		c.Debug = true
	}
	return
}

func (c *Config) loadConfig() (err error) {
	var data []byte
	if data, err = ioutil.ReadFile(c.configFile); err != nil {
		return
	}
	return c.unmarshal(data, c)
}

func (c *Config) GetEnvName() Env {
	return c.EnvName
}

func (c *Config) DebugEnabled() bool {
	return c.Debug
}

func (c *Config) unmarshal(data []byte, i interface{}) (err error) {
	return yaml.Unmarshal(data, i)
}
