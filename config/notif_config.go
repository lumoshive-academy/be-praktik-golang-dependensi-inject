package config

// Config adalah struct untuk konfigurasi aplikasi
type NotifConfig struct {
	AppName string
}

// NewConfig menyediakan instance dari Config
func NewNotifConfig() *NotifConfig {
	return &NotifConfig{
		AppName: "NotificationApp",
	}
}
