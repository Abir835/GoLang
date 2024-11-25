package main

//
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
//EmailNotification struct
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
//// SMSNotification struct
//type SMSNotification struct{}
//
//func NewSMSNotification() NotificationSender {
//	return &SMSNotification{}
//}
//
//func (s *SMSNotification) Send(message string) error {
//	fmt.Printf("Sending SMS: %s\n", message)
//	return nil
//}
//
//// PushNotification struct
//type PushNotification struct{}
//
//func NewPushNotification() NotificationSender {
//	return &PushNotification{}
//}
//
//func (p *PushNotification) Send(message string) error {
//	fmt.Printf("Sending push notification: %s\n", message)
//	return nil
//}
//
//// NotificationType enum
//type NotificationType int
//
//const (
//	Email NotificationType = iota
//	SMS
//	Push
//)
//
//// NotificationFactory for creating notification senders
//type NotificationFactory struct{}
//
//func (f *NotificationFactory) CreateNotificationSender(nType NotificationType, args ...interface{}) NotificationSender {
//	switch nType {
//	case Email:
//		if len(args) < 6 {
//			return nil
//		}
//		return NewEmailNotification(
//			args[0].(string),   // From
//			args[1].([]string), // To
//			args[2].(string),   // Subject
//			args[3].(string),   // SMTPHost
//			args[4].(string),   // SMTPPort
//			args[5].(string),   // Username
//			args[6].(string),   // Password
//		)
//	case SMS:
//		return NewSMSNotification()
//	case Push:
//		return NewPushNotification()
//	default:
//		return nil
//	}
//}
//
//// NotificationStrategy struct
//type NotificationStrategy struct {
//	sender NotificationSender
//}
//
//func NewNotificationStrategy(factory *NotificationFactory, nType NotificationType, args ...interface{}) *NotificationStrategy {
//	return &NotificationStrategy{
//		sender: factory.CreateNotificationSender(nType, args...),
//	}
//}
//
//func (s *NotificationStrategy) SendNotification(message string) error {
//	return s.sender.Send(message)
//}
//
//// NotificationService struct
//type NotificationService struct {
//	strategy *NotificationStrategy
//}
//
//func NewNotificationService(strategy *NotificationStrategy) *NotificationService {
//	return &NotificationService{
//		strategy: strategy,
//	}
//}
//
//func (s *NotificationService) Notify(message string) error {
//	return s.strategy.SendNotification(message)
//}
//
//func main() {
//	factory := &NotificationFactory{}
//
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
//	emailStrategy := NewNotificationStrategy(factory, Email, from, to, subject, smtpHost, smtpPort, username, password)
//	emailService := NewNotificationService(emailStrategy)
//	if err := emailService.Notify("Hello via email"); err != nil {
//		fmt.Println("Error:", err)
//	}
//
//	// Create SMS notification service
//	smsStrategy := NewNotificationStrategy(factory, SMS)
//	smsService := NewNotificationService(smsStrategy)
//	smsService.Notify("Hello via SMS")
//
//	// Create push notification service
//	pushStrategy := NewNotificationStrategy(factory, Push)
//	pushService := NewNotificationService(pushStrategy)
//	pushService.Notify("Hello via push notification")
//}
