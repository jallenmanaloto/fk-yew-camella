package main

import (
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

	// Do not run the app
	if !config.Enable {
		return
	}

	// Generate github workflow to run the app on schedule
	cron := config.CronExpression()
	generator.ScheduledWorkflow(cron)

	// Initialize mailer and send the email
	m := mailer.New(
		"smtp.gmail.com",
		"587",
		config.Email,
		config.Password,
	)
	m.Send(config)
}
