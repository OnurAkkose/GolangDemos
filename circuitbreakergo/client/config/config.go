package config

type ConfigurationApiConfigDatabase struct {
	Host    string `yaml:"host" env:"CONFIGURATION_API_HOST" env-default:"localhost"`
	Port    int    `yaml:"port" env:"CONFIGURATION_API_PORT" env-default:"6161"`
	Version int    `yaml:"version" env:"CONFIGURATION_API_VERSION" env-default:"1"`
}
