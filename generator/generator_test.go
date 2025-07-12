package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestContentIncludesCron(t *testing.T) {
	cron := "0 0 * * *"
	result := content(cron)

	if !strings.Contains(result, cron) {
		t.Errorf("Expects cron string '%s' to be in the workflow, but not found", cron)
	}
}

func TestScheduledWorkflow(t *testing.T) {
	tmpDir := t.TempDir()
	testPath := filepath.Join(tmpDir, "scheduler.yml")
	cron := "0 0 * * *"

	ScheduledWorkflow(cron, testPath)

	// Assert file exists
	content, err := os.ReadFile(testPath)
	if err != nil {
		t.Fatalf("Expected file '%s' to exist, but got error: %v", testPath, err)
	}

	if !strings.Contains(string(content), cron) {
		t.Fatalf("Expected file content to include cron '%s', but got:\n'%s'", cron, string(content))
	}
}
