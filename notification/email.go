package notification

import "fmt"

// EmailService adalah struct untuk layanan email
type EmailService struct{}

// SendEmail mengirim email
func (e *EmailService) SendEmail(message string) {
	fmt.Println("Sending email:", message)
}

// NewEmailService menyediakan instance dari EmailService
func NewEmailService() *EmailService {
	return &EmailService{}
}
