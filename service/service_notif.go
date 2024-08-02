package service

import (
	"golang-dependensi-inject/config"
	"golang-dependensi-inject/notification"
)

// Notifier adalah struct yang menggunakan EmailService, SMSService, dan Config
type Notifier struct {
	EmailService *notification.EmailService
	SMSService   *notification.SMSService
	NotifConfig  *config.NotifConfig
}

// SendNotification mengirim notifikasi melalui email dan SMS
func (n *Notifier) SendNotification(message string) {
	n.EmailService.SendEmail(message)
	n.SMSService.SendSMS(message)
}
