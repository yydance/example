package storage

type WatchResponse struct {
	Events []Event
	Error  error
}

type Event struct {
	KeyPair
	Type EventType
}

type KeyPair struct {
	Key   string
	Value string
}

type EventType string

var (
	EventTypePut    EventType = "put"
	EventTypeDelete EventType = "delete"
)
