package events

import "time"

// DeviceCreatedEvent represents a device creation event
type DeviceCreatedEvent struct {
	EventID    string     `json:"event_id"`
	EventName  string     `json:"event_name"`
	EventTime  time.Time  `json:"event_time"`
	DeviceID   string     `json:"device_id"`
	AccountID  string     `json:"account_id"`
	DeviceInfo DeviceInfo `json:"device_info"`
}

// DeviceInfo contains device information
type DeviceInfo struct {
	DeviceName   string `json:"device_name"`
	DeviceType   string `json:"device_type"`
	SerialNumber string `json:"serial_number"`
}

// DeviceDeletedEvent represents a device deletion event
type DeviceDeletedEvent struct {
	EventID   string    `json:"event_id"`
	EventName string    `json:"event_name"`
	EventTime time.Time `json:"event_time"`
	DeviceID  string    `json:"device_id"`
	AccountID string    `json:"account_id"`
}

// DeviceUpdatedEvent represents a device update event
type DeviceUpdatedEvent struct {
	EventID    string     `json:"event_id"`
	EventName  string     `json:"event_name"`
	EventTime  time.Time  `json:"event_time"`
	DeviceID   string     `json:"device_id"`
	AccountID  string     `json:"account_id"`
	DeviceInfo DeviceInfo `json:"device_info"`
	Changes    []string   `json:"changes"`
}
