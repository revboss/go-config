package config_test

import (
	"github.com/revboss/go-config"
	"os"
	"testing"
)

type Config struct {
	String string
	Int    int
}

func TestConfigEnvString(t *testing.T) {
	cc := &Config{}

	c := &config.Config{
		Env: map[string]config.Value{
			"STRING": {
				Optional: false,
				Value:    &cc.String,
			},
		},
	}

	os.Setenv("STRING", "test")

	e := c.Load()
	if e != nil {
		t.Error(e)
	}

	if cc.String != "test" {
		t.Errorf("Expected %q, got %q", "test", cc.String)
	}

	os.Unsetenv("STRING")
}

func TestConfigEnvStringFail(t *testing.T) {
	cc := &Config{}
	c := &config.Config{
		Env: map[string]config.Value{
			"STRING": {
				Optional: false,
				Value:    &cc.String,
			},
		},
	}

	e := c.Load()
	if e == nil {
		t.Errorf("Load should have failed since environment variable BLAH is not optional")
	}
}

func TestConfigEnvStringOptional(t *testing.T) {
	cc := &Config{}
	c := &config.Config{
		Env: map[string]config.Value{
			"STRING": {
				Optional: true,
				Value:    &cc.String,
			},
		},
	}

	e := c.Load()
	if e != nil {
		t.Errorf("Load should not have failed since environment variable BLAH is optional")
	}
}

func TestConfigEnvInt(t *testing.T) {
	cc := &Config{}
	c := &config.Config{
		Env: map[string]config.Value{
			"INT": {
				Optional: false,
				Value:    &cc.Int,
			},
		},
	}

	os.Setenv("INT", "5")

	e := c.Load()
	if e != nil {
		t.Error(e)
	}

	if cc.Int != 5 {
		t.Errorf("Expected %d, got %d", 5, cc.Int)
	}

	os.Unsetenv("INT")
}

func TestConfigEnvIntFail(t *testing.T) {
	cc := &Config{}
	c := &config.Config{
		Env: map[string]config.Value{
			"INT": {
				Optional: false,
				Value:    &cc.Int,
			},
		},
	}

	e := c.Load()
	if e == nil {
		t.Errorf("Load should have failed since environment variable BLAH is not optional")
	}
}

func TestConfigEnvIntOptional(t *testing.T) {
	cc := &Config{}
	c := &config.Config{
		Env: map[string]config.Value{
			"INT": {
				Optional: true,
				Value:    &cc.Int,
			},
		},
	}

	e := c.Load()
	if e != nil {
		t.Errorf("Load should not have failed since environment variable BLAH is optional")
	}
}

func TestLoadConfig(t *testing.T) {
	cc := &Config{}

	os.Setenv("STRING", "test")

	e := config.LoadConfig(config.Config{
		Env: map[string]config.Value{
			"STRING": {
				Optional: false,
				Value:    &cc.String,
			},
			"INT": {
				Optional: true,
				Value:    &cc.Int,
			},
		},
	})
	if e != nil {
		t.Error("LoadConfig should not fail since STRING is defined and INT is optional")
	}

	os.Unsetenv("STRING")
}
