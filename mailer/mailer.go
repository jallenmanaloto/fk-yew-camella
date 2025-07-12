package mailer

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/fk-yew-camella/config"
)

type Mailer struct {
	Host     string
	Port     string
	Username string
	Password string
}

func New(host, port, username, password string) *Mailer {
	return &Mailer{
		host, port, username, password,
	}
}

func (m *Mailer) Send(config *config.Config) error {
	addr := fmt.Sprintf("%s:%s", m.Host, m.Port)

	var recepients []string
	recepients = append(recepients, config.To...)
	recepients = append(recepients, config.Bcc...)
	recepients = append(recepients, config.Cc...)

	// Format headers
	headers := m.formatHeaders(
		config.To,
		config.Cc,
		config.Bcc,
		config.Subject,
	)

	// Build email
	var msg strings.Builder
	for k, v := range headers {
		fmt.Fprintf(&msg, "%s: %s\r\n", k, v)
	}
	msg.WriteString("\r\n")
	msg.WriteString(config.MessageBody)

	// Authenticate and send email
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)
	if err := smtp.SendMail(addr, auth, m.Username, recepients, []byte(msg.String())); err != nil {
		return fmt.Errorf("Failed to send email: %w", err)
	}

	return nil
}

func (m *Mailer) formatHeaders(to, cc, bcc []string, subject string) map[string]string {
	headers := make(map[string]string)
	headers["From"] = m.Username
	headers["Subject"] = subject
	headers["To"] = strings.Join(to, ", ")
	if len(cc) > 0 {
		headers["Cc"] = strings.Join(cc, ", ")
	}

	if len(bcc) > 0 {
		headers["Bcc"] = strings.Join(bcc, ", ")
	}

	return headers
}
