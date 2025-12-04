package events

import "time"

// ClientKVMSessionRequestedEvent represents a client KVM session request event
type ClientKVMSessionRequestedEvent struct {
    EventID   string    `json:"event_id"`
    EventName string    `json:"event_name"`
    EventTime time.Time `json:"event_time"`
    ClientID  string    `json:"client_id"`
    DeviceID  string    `json:"device_id"`
    AccountID string    `json:"account_id"`
    //SessionID string    `json:"session_id"`
}

// ClientKVMSessionCloseEvent represents a client KVM session close event
type ClientKVMSessionCloseEvent struct {
    EventID   string    `json:"event_id"`
    EventName string    `json:"event_name"`
    EventTime time.Time `json:"event_time"`
    ClientID  string    `json:"client_id"`
    DeviceID  string    `json:"device_id"`
    SessionID string    `json:"session_id"`
    Duration  int64     `json:"duration"` // Duration in seconds
}
