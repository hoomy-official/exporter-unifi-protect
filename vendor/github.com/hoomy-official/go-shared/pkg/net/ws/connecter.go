package ws

import (
	"context"
	"io"
)

// MessageType specifies the type of WebSocket messages.
type MessageType int

// Possible message types as per RFC 6455.
const (
	MessageText   MessageType = iota + 1 // Textual data message
	MessageBinary                        // Binary data message
)

// StatusCode defines status codes used to close WebSocket connections.
type StatusCode int

// Connecter allows interaction with a WebSocket connection.
type Connecter interface {
	// Reader starts reading a message from the connection.
	Reader(ctx context.Context) (MessageType, io.Reader, error)

	// Read retrieves a complete message from the connection.
	Read(ctx context.Context) (MessageType, []byte, error)

	// Close sends a close message and waits for acknowledgment.
	Close(code StatusCode, reason string) error

	// CloseNow forcefully closes the connection without a closing handshake.
	CloseNow() error
}
