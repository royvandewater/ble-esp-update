package ring

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/darwin"
)

var _ServiceUUID = ble.MustParse("2ba75e8a-5b5b-447b-ab9a-b79e21dd64e0")
var _ColorCharacteristicUUID = ble.MustParse("08f490bf-28f1-4d55-897d-ab8d74effffb")
var _CommandCharacteristicUUID = ble.MustParse("04b29961-90fd-4ee7-bb48-f203bde84f44")

type _BLERing struct {
	device                ble.Device
	localName             string
	client                ble.Client
	commandCharacteristic *ble.Characteristic
	colorCharacteristic   *ble.Characteristic
}

func _NewBLERing(localName string) (*_BLERing, error) {
	device, err := darwin.NewDevice()
	if err != nil {
		return nil, err
	}
	ble.SetDefaultDevice(device)

	return &_BLERing{client: nil, device: device, localName: localName}, nil
}

func (ring *_BLERing) Connect(timeout time.Duration) error {
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), timeout))

	client, err := ble.Connect(ctx, ring.localNameFilter)
	if err != nil {
		return err
	}
	services, err := client.DiscoverServices([]ble.UUID{_ServiceUUID})
	if err != nil {
		return err
	}
	if len(services) != 1 {
		return fmt.Errorf("Expected exactly 1 service, got: %v", services)
	}

	characteristics, err := client.DiscoverCharacteristics([]ble.UUID{_ColorCharacteristicUUID, _CommandCharacteristicUUID}, services[0])
	if err != nil {
		return err
	}
	if len(characteristics) != 2 {
		return fmt.Errorf("Expected exactly 2 servics, got: %v", characteristics)
	}

	for _, characteristic := range characteristics {
		if characteristic.UUID.Equal(_ColorCharacteristicUUID) {
			ring.colorCharacteristic = characteristic
		}
		if characteristic.UUID.Equal(_CommandCharacteristicUUID) {
			ring.commandCharacteristic = characteristic
		}
	}

	ring.client = client
	return nil
}

func (ring *_BLERing) Disconnect() error {
	if ring.client == nil {
		return nil
	}

	return ring.client.CancelConnection()
}

func (ring *_BLERing) PulseColor(r, g, b byte) error {
	if ring.client == nil {
		return nil
	}

	err := ring.client.WriteCharacteristic(ring.colorCharacteristic, []byte{r, g, b}, false)
	if err != nil {
		return err
	}

	return ring.client.WriteCharacteristic(ring.commandCharacteristic, []byte{1}, false)
}

func (ring *_BLERing) SetColor(r, g, b byte) error {
	if ring.client == nil {
		return nil
	}

	err := ring.client.WriteCharacteristic(ring.colorCharacteristic, []byte{r, g, b}, false)
	if err != nil {
		return err
	}

	return ring.client.WriteCharacteristic(ring.commandCharacteristic, []byte{0}, false)
}

func (ring *_BLERing) localNameFilter(a ble.Advertisement) bool {
	debug("Found BLE Device %v", a.LocalName())
	return strings.ToUpper(a.LocalName()) == strings.ToUpper(ring.localName)
}
