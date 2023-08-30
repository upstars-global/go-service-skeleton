package config

type LoggerConfigProvider interface {
	GetLoggerFormat() string
	GetLoggerLevel() string
}

type logger struct {
	Format string `yaml:"format"`
	Level  string `yaml:"level"`
}

func (c *Config) GetLoggerFormat() string {
	return c.Logger.Format
}

func (c *Config) GetLoggerLevel() string {
	return c.Logger.Level
}
