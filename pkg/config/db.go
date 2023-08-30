package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type db struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Name            string `yaml:"name"`
	Migrations      string `yaml:"migrations"`
	SSLModeDisabled bool   `yaml:"sslModeDisabled"`
	MaxOpenCons     int    `yaml:"maxOpenCons"`
	MaxIdleCons     int    `yaml:"maxIdleCons"`
}

type DBConfigProvider interface {
	GetDBDSN() string
	GetDBMigrations() string
	GetDBMaxIdleCons() int
	GetDBMaxOpenCons() int
}

func (c Config) GetDBDSN() string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		c.DB.Host,
		c.DB.Port,
		c.DB.Username,
		c.DB.Password,
		c.DB.Name,
	)
	if c.DB.SSLModeDisabled {
		dsn += " sslmode=disable"
	}
	return dsn
}

func (c Config) GetDBMigrations() string {
	return c.DB.Migrations
}

func (c Config) GetDBMaxIdleCons() int {
	return c.DB.MaxIdleCons
}

func (c Config) GetDBMaxOpenCons() int {
	return c.DB.MaxOpenCons
}

func (c *Config) OverrideDBConfig() {
	if viper.IsSet("db-host") {
		c.DB.Host = viper.GetString("db-host")
	}
	if viper.IsSet("db-username") {
		c.DB.Username = viper.GetString("db-username")
	}
	if viper.IsSet("db-password") {
		c.DB.Password = viper.GetString("db-password")
	}
	if viper.IsSet("db-name") {
		c.DB.Name = viper.GetString("db-name")
	}
	if viper.IsSet("db-migrations") {
		c.DB.Migrations = viper.GetString("db-migrations")
	}

	if viper.IsSet("db-port") {
		c.DB.Port = viper.GetInt("db-port")
	}
	if viper.IsSet("db-max-open-cons") {
		c.DB.MaxOpenCons = viper.GetInt("db-max-open-cons")
	}
	if viper.IsSet("db-max-idle-cons") {
		c.DB.MaxIdleCons = viper.GetInt("db-max-idle-cons")
	}

	if viper.IsSet("db-ssl-mode-disabled") {
		c.DB.SSLModeDisabled = viper.GetBool("db-ssl-mode-disabled")
	}
}
