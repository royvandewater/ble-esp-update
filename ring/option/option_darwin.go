package option

import "github.com/paypal/gatt"

// DefaultClientOptions define the default options for
// creating a new client
var DefaultClientOptions = []gatt.Option{
	gatt.MacDeviceRole(gatt.CentralManager),
}
