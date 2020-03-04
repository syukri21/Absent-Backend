package socket

// Message ...
type Message struct {
	Type string       `json:"type"`
	Data *interface{} `json:"data"`
}

const (
	// NewGenerateQrcode ...
	NewGenerateQrcode MessageType = "GENERATE_QRCODE"
)

// MessageType ...
type MessageType string
