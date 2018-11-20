package config

import (
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config is the top-level configuration for dingtalk
type Config struct {
	Profiles map[string]string `yaml:"profiles,omitempty"`
}

// Load parses the YAML input s into a Config.
func Load(s string) (*Config, error) {
	cfg := &Config{}
	err := yaml.Unmarshal([]byte(s), cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Profiles == nil {
		return nil, errors.New("no ding profiles provided in config")
	}

	for profile, webhookURL := range cfg.Profiles {
		profile = strings.TrimSpace(profile)
		webhookURL = strings.TrimSpace(webhookURL)
		if profile == "" {
			return nil, errors.New("profile part cannot be empty")
		}
		if webhookURL == "" {
			return nil, errors.New("webhook-url part cannot be emtpy")
		}
	}

	return cfg, nil
}

// LoadFile parses the given YAML file into a Config.
func LoadFile(filename string) (*Config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cfg, err := Load(string(content))
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
