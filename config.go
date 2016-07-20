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
	Default  interface{}
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
		vp := v.Value.(*int)
		i, _ := strconv.Atoi(s)
		if i != 0 {
			*vp = i
		} else if v.Default != nil {
			*vp = v.Default.(int)
		}

		return *vp != 0 || v.Optional

	case *string:
		vp := v.Value.(*string)
		if len(s) > 0 {
			*vp = s
		} else if v.Default != nil {
			*vp = v.Default.(string)
		}

		return len(*vp) > 0 || v.Optional
	}

	return v.Optional
}
