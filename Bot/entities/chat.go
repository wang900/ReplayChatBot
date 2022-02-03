package entities

type Chat struct {
	ID           int           `json:"id"`
	HostID       string        `json:"host_id"`
	ChatMessages []ChatMessage `json:"chat_messages"`
}
