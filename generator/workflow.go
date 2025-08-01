package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func ScheduledWorkflow(cron, path string) {
	c := content(cron)

	// Write content to .github/workflows/scheduler.yml file
	// Panic since this will not run the app on the defined schedule preference
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, []byte(c), 0644); err != nil {
		panic(err)
	}
}

func content(cron string) string {
	workflow := fmt.Sprintf(`name: Scheduled Run

on:
  schedule:
    - cron: '%s'

jobs:
  run-app:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: Setup app
        uses: actions/setup-go@v5
        with:
          go-version: '^1.23'

      - name: Build app
        run: go build -v -o mailer

      - name: Run app
        run: ./mailer
`, cron)

	return workflow
}
