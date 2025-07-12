package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fk-yew-camella/config"
	"github.com/fk-yew-camella/generator"
	"github.com/fk-yew-camella/mailer"
)

const CONFIG_PATH = "editme.json"

func main() {
	config, err := config.Load(CONFIG_PATH)
	if err != nil {
		panic(err) // what's the use if we can't load the config file, right?
	}

	// Generate github workflow to run the app on schedule
	cron := config.CronExpression()
	targetPath := filepath.Join(".github", "workflows", "scheduler.yml")
	generator.ScheduledWorkflow(cron, targetPath)

	// Check for --generate flag
	if len(os.Args) > 1 && os.Args[1] == "--generate" {
		fmt.Println("✅ Workflow file generated at .github/workflows/scheduler.yml")
		fmt.Println("👉 Commit and push it to activate the scheduled job.")
		return
	}

	// We feel like a saint, so we won't run the app
	if !config.Enable {
		return
	}

	// Initialize mailer and send the email
	m := mailer.New(
		"smtp.gmail.com",
		"587",
		config.Email,
		config.Password,
	)
	m.Send(config)
}
