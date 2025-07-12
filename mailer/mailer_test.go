package mailer

import (
	"errors"
	"net/smtp"
	"strings"
	"testing"

	"github.com/fk-yew-camella/config"
)

func TestSend_EmailBuildAndSuccess(t *testing.T) {
	mockCalled := false
	mockSend := func(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
		mockCalled = true

		content := string(msg)

		// Assert headers are included
		if !strings.Contains(content, "Subject: Test Email") {
			t.Errorf("Missing subject header")
		}
		if !strings.Contains(content, "To: test@email.com, fker@email.com") {
			t.Errorf("Missing or wrong 'To' header")
		}
		if !strings.Contains(content, "Cc: cc@email.com") {
			t.Errorf("Missing or wrong 'Cc' header")
		}
		if !strings.Contains(content, "Bcc: schlong@email.com") {
			t.Errorf("Missing or wrong 'Bcc' header")
		}

		// Assert body message
		if !strings.Contains(content, "This is the body of the message") {
			t.Errorf("Missing body message")
		}

		return nil
	}

	m := New("smpt.example.com", "587", "user@example.com", "justapassword")
	m.sendEmail = mockSend

	cfg := &config.Config{
		To:          []string{"test@email.com", "fker@email.com"},
		Cc:          []string{"cc@email.com"},
		Bcc:         []string{"schlong@email.com"},
		Subject:     "Test Email",
		MessageBody: "This is the body of the message",
	}

	err := m.Send(cfg)
	if err != nil {
		t.Fatalf("Expected no error sending, got: %v", err)
	}

	if !mockCalled {
		t.Fatalf("Expected mockSend to be called, but wasn't")
	}
}

func TestSend_ErrorIfSendFailes(t *testing.T) {
	mockSend := func(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
		return errors.New("send failure")
	}

	m := New("smpt.example.com", "587", "user@example.com", "justapassword")
	m.sendEmail = mockSend

	cfg := &config.Config{
		To:          []string{"target@email.com"},
		Subject:     "Should Fail Test",
		MessageBody: "This email should fail coz it should",
	}

	err := m.Send(cfg)
	if err == nil {
		t.Fatalf("Expecting an error, but got nil")
	}

	if !strings.Contains(err.Error(), "send failure") {
		t.Errorf("Unexpected error received: %v", err)
	}
}
