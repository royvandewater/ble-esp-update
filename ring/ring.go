package ring

import "time"
import DEBUG "github.com/computes/go-debug"

var debug = DEBUG.Debug("ble-build-status:ring")

// Ring is the interface for managing an LED ring
type Ring interface {
	// Connect establishes a connection with the given timeout
	Connect(timeout time.Duration) error

	// Disconnect disconnects
	Disconnect() error

	// PulseColor sets the color and pulses
	PulseColor(r, g, b byte) error

	// SetColor sets the color
	SetColor(r, g, b byte) error
}

// New constructs an unconnected ring instance
func New(localName string) (Ring, error) {
	return _NewBLERing(localName)
}
