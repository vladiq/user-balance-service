package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

var cfg *Config

func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}
	return Config{}
}

type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
}

type Server struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	BasePath string `yaml:"basePath"`
}

// DB - contains db information
type DB struct {
	DSN             string        `yaml:"DSN"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"MaxIdleConns"`
	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

func (db *DB) GetDSN() string {
	return db.DSN
}

func (db *DB) GetMaxOpenConns() int {
	return db.MaxOpenConns
}

func (db *DB) GetMaxIdleConns() int {
	return db.MaxIdleConns
}

func (db *DB) GetConnMaxIdleTime() time.Duration {
	return db.ConnMaxIdleTime
}

func (db *DB) GetConnMaxLifetime() time.Duration {
	return db.ConnMaxLifetime
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project Project `yaml:"project"`
	Server  Server  `yaml:"server"`
	DB      DB      `yaml:"db"`
}

// ReadConfigYML - read config from a file and initialize a Config instance.
func ReadConfigYML(configYML string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(configYML)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	return nil
}
