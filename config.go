package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Env map[string]Value
}

type Value struct {
	Optional bool
	Value    interface{}
}

func LoadConfig(config Config) error {
	return (&config).Load()
}

func (c *Config) Load() error {
	return c.loadEnv()
}

func (c *Config) loadEnv() error {
	for k, v := range c.Env {
		env := os.Getenv(k)
		if !set(&v, env) {
			return fmt.Errorf("Missing environment variable: %s", k)
		}
	}
	return nil
}

func set(v *Value, s string) bool {
	switch v.Value.(type) {
	case *int:
		i, _ := strconv.Atoi(s)
		if i != 0 || v.Optional == true {
			*v.Value.(*int) = i
			return true
		}

	case *string:
		if len(s) > 0 || v.Optional == true {
			*v.Value.(*string) = s
			return true
		}
	}
	return false
}
