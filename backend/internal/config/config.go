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
	Redis struct {
		Host     string `yaml:"host"`
		Port     uint   `yaml:"port"`
		DB       int    `yaml:"db"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Logger struct {
		Format string `yaml:"format"`
		Level  string `yaml:"level"`
		Path   string `yaml:"path"`
	} `yaml:"logger"`
	Jwt struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}
