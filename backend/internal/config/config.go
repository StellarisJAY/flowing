package config

import "time"

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
	FileStore string `yaml:"fileStore"`
	Minio     struct {
		Endpoint  string `yaml:"endpoint"`
		AccessKey string `yaml:"accessKey"`
		SecretKey string `yaml:"secretKey"`
		Bucket    string `yaml:"bucket"`
	} `yaml:"minio"`
	Worker string `yaml:"worker"`
	GoPool struct {
		Size    int           `yaml:"size"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"gopool"`
}
