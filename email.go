package main

//
//import (
//	"fmt"
//	"net/smtp"
//)
//
//// NotificationSender interface
//type NotificationSender interface {
//	Send(message string) error
//}
//
//// EmailNotification struct
//type EmailNotification struct {
//	From     string
//	To       []string
//	Subject  string
//	SMTPHost string
//	SMTPPort string
//	Username string
//	Password string
//}
//
//func NewEmailNotification(from string, to []string, subject string, smtpHost string, smtpPort string, username string, password string) NotificationSender {
//	return &EmailNotification{
//		From:     from,
//		To:       to,
//		Subject:  subject,
//		SMTPHost: smtpHost,
//		SMTPPort: smtpPort,
//		Username: username,
//		Password: password,
//	}
//}
//
//func (e *EmailNotification) Send(message string) error {
//	// Set up authentication information.
//	auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPHost)
//
//	// Compose the email.
//	body := fmt.Sprintf("Subject: %s\r\n\r\n%s", e.Subject, message)
//
//	// Connect to the SMTP server.
//	addr := fmt.Sprintf("%s:%s", e.SMTPHost, e.SMTPPort)
//	if err := smtp.SendMail(addr, auth, e.From, e.To, []byte(body)); err != nil {
//		return fmt.Errorf("failed to send email: %w", err)
//	}
//
//	fmt.Printf("Email sent to: %s\n", e.To)
//	return nil
//}
//
//func main() {
//	// SMTP configuration
//	from := "your_email@example.com"
//	to := []string{"recipient@example.com"}
//	subject := "Test Email"
//	smtpHost := "smtp.example.com" // Your SMTP host
//	smtpPort := "587"              // Your SMTP port
//	username := "your_email@example.com"
//	password := "your_email_password"
//
//	// Create email notification service
//	emailService := NewEmailNotification(from, to, subject, smtpHost, smtpPort, username, password)
//	if err := emailService.Send("Hello via email"); err != nil {
//		fmt.Println("Error:", err)
//	}
//}
