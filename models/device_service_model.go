package models

type DeviceServices struct {
	ID        int    `json:"id"`
	DeviceID  string `json:"device_id"`
	ServiceID string `json:"service_id"`
}
