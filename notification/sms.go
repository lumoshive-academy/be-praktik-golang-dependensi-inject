package notification

import "fmt"

// SMSService adalah struct untuk layanan SMS
type SMSService struct{}

// SendSMS mengirim SMS
func (s *SMSService) SendSMS(message string) {
	fmt.Println("Sending SMS:", message)
}

// NewSMSService menyediakan instance dari SMSService
func NewSMSService() *SMSService {
	return &SMSService{}
}
