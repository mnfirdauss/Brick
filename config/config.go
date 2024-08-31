package config

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

/* Environment utility */

func loadEnvStr(key string, result *string) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return
	}

	*result = s
}

func loadEnvUint(key string, result *uint) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return
	}

	n, err := strconv.Atoi(s)

	if err != nil {
		return
	}

	*result = uint(n)
}

/* Configuration */
type listenConfig struct {
	Host string `yaml:"host" json:"host"`
	Port uint   `yaml:"port" json:"port"`
}

func (l listenConfig) Addr() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

func defaultListenConfig() listenConfig {
	return listenConfig{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

func (l *listenConfig) loadFromEnv() {
	loadEnvStr("LISTEN_HOST", &l.Host)
	loadEnvUint("LISTEN_PORT", &l.Port)
}

type pgConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     uint   `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`

	DBName  string `yaml:"db_name" json:"db_name"`
	SslMode string `yaml:"ssl_mode" json:"ssl_mode"`
}

func (p pgConfig) ConnStr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", p.Username, p.Password, p.Host, p.Port, p.DBName)
}

func defaultPgConfig() pgConfig {
	return pgConfig{
		Host:    "localhost",
		Port:    5432,
		DBName:  "todo",
		SslMode: "disable",
	}
}

func (p *pgConfig) loadFromEnv() {
	loadEnvStr("DB_HOST", &p.Host)
	loadEnvUint("DB_PORT", &p.Port)
	loadEnvStr("DB_USERNAME", &p.Username)
	loadEnvStr("DB_PASSWORD", &p.Password)
	loadEnvStr("DB_NAME", &p.DBName)
	loadEnvStr("DB_SSL", &p.SslMode)

}

type baseURL struct {
	BankURL string `yaml:"bank_url" json:"bank_url"`
}

func (l *baseURL) loadFromEnv() {
	loadEnvStr("BANK_BASE_URL", &l.BankURL)
}

type config struct {
	Listen   listenConfig `yaml:"listen" json:"listen"`
	BaseURL  baseURL      `yaml:"base_url" json:"base_url"`
	DBConfig pgConfig     `yaml:"db" json:"db"`
}

func (c *config) LoadFromEnv() {
	c.Listen.loadFromEnv()
	c.BaseURL.loadFromEnv()
	c.DBConfig.loadFromEnv()
}

func DefaultConfig() config {
	return config{
		Listen:   defaultListenConfig(),
		DBConfig: defaultPgConfig(),
	}
}

func loadConfigFromReader(r io.Reader, c *config) error {
	return yaml.NewDecoder(r).Decode(c)
}

func LoadConfigFromFile(fn string, c *config) error {
	_, err := os.Stat(fn)

	if err != nil {
		return err
	}

	f, err := os.Open(fn)

	if err != nil {
		return err
	}

	defer f.Close()

	return loadConfigFromReader(f, c)
}
