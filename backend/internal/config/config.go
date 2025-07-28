package config

type Config struct {
	DB struct {
		Driver string `yaml:"driver"`
	} `yaml:"db"`
	Postgres struct {
		DSN string `yaml:"dsn"`
	} `yaml:"postgres"`
	MySQL struct {
		DSN string `yaml:"dsn"`
	} `yaml:"mysql"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}
