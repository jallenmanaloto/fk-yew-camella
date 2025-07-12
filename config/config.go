package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Bcc         []string `json:"bcc,omitempty"`
	Cc          []string `json:"cc,omitempty"`
	Cron        string   `json:"cron,omitempty"`
	Email       string   `json:"email"`
	Enable      bool     `json:"enable"`
	MessageBody string   `json:"message_body"`
	Password    string   `json:"password"`
	Schedule    string   `json:"schedule"`
	Subject     string   `json:"subject"`
	To          []string `json:"to"`
}

const CONFIG_PATH = "editme.json"

func Load() (*Config, error) {
	file, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file from path: %s, err: %w", CONFIG_PATH, err)
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("Failed to parse file from path: %s, err: %w", CONFIG_PATH, err)
	}

	return &config, nil
}

func (c *Config) CronExpression() string {
	switch c.Schedule {
	case "daily":
		return "0 0 * * *"
	case "weekly":
		return "0 0 * * 0"
	case "hourly":
		return "0 * * * *"
	case "custom":
		return c.Cron
	default:
		return "0 0 * * *" // fallback to daily schedule
	}
}
