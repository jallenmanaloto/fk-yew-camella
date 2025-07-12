package config

import (
	"os"
	"testing"
)

const TEMP_CONFIG_PATH = "temp.json"

func writeTempConfigFile(t *testing.T, content string) {
	t.Helper()
	err := os.WriteFile(TEMP_CONFIG_PATH, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write to temp config file: %v", err)
	}
}

func cleanupTempConfigFile(t *testing.T) {
	t.Helper()
	if err := os.Remove(TEMP_CONFIG_PATH); err != nil {
		t.Fatalf("Failed to cleanup temp config file: %v", err)
	}
}

func TestLoad(t *testing.T) {
	tmpFile := "test_config.json"
	t.Cleanup(func() { os.Remove(tmpFile) })

	err := os.WriteFile(tmpFile, []byte(`{
		"email": "me@example.com",
		"password": "app-pass",
		"message_body": "Hello!",
		"to": ["you@example.com"],
		"subject": "Test",
		"schedule": "daily",
		"enable": true
		}`), 0644)
	if err != nil {
		t.Fatalf("Failed to write test config file: %v", err)
	}

	cfg, err := Load(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load config file: %v", err)
	}

	if cfg.Email != "me@example.com" {
		t.Fatalf("Expected email to be me@example.com, but got %s", cfg.Email)
	}
}

func TestCronExpression(t *testing.T) {
	tests := []struct {
		schedule string
		cron     string
		custom   string
	}{
		{"daily", "0 0 * * *", ""},
		{"weekly", "0 0 * * 0", ""},
		{"hourly", "0 * * * *", ""},
		{"custom", "42 12 * * *", "42 12 * * *"},
		{"invalid", "0 0 * * *", ""},
	}

	for _, test := range tests {
		c := &Config{
			Schedule: test.schedule,
			Cron:     test.cron,
		}

		got := c.CronExpression()
		if got != test.cron {
			t.Errorf("Schedule '%s': expected '%s', but got '%s'", test.schedule, test.cron, got)
		}
	}
}
