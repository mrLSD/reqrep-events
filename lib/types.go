package lib

// EventProxy - Event server data
type EventProxy struct {
	Url   string
	Name  string
	Event EventData
}

// EventData - data for event sending
type EventData struct {
	EventName string      `json:"eventName"`
	EventId   int64       `json:"eventId"`
	Value     interface{} `json:"value"`
}

// EventHandler- Handler for Event Server response
type EventHandler func(ev EventData) (EventData, error)
