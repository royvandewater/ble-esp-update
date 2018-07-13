package main

import (
	"fmt"
	"log"
	"time"

	"github.com/royvandewater/ble-esp-update/ring"
)

var color = map[string][]byte{
	"black":       []byte{0x00, 0x00, 0x00},
	"deep-purple": []byte{0x20, 0x00, 0x90},
	"light-pink":  []byte{0xff, 0x40, 0x40},
	"gray":        []byte{0x10, 0x10, 0x10},
	"green":       []byte{0x00, 0xff, 0x00},
	"orange":      []byte{0xff, 0x20, 0x00},
	"pink":        []byte{0xff, 0x20, 0x20},
	"purple":      []byte{0x40, 0x00, 0xff},
	"white":       []byte{0xff, 0xff, 0xff},
	"yellow":      []byte{0xff, 0xff, 0x00},
}["orange"]

func fatalIfErrorf(err error, msg string, rest ...interface{}) {
	if err == nil {
		return
	}

	log.Fatalln(fmt.Sprintf(msg, rest...), err.Error())
}

func main() {
	ringName := "esp32-neopixel"

	r, err := ring.New(ringName)
	fatalIfErrorf(err, "Failed to construct a new ring")

	err = r.Connect(10 * time.Second)
	fatalIfErrorf(err, "Failed to connect to ring")

	err = r.SetColor(color[0], color[1], color[2])
	fatalIfErrorf(err, "Failed to set color on ring")

	<-time.After(2 * time.Second)
	err = r.PulseColor(color[0], color[1], color[2])
	fatalIfErrorf(err, "Failed to pulse color on ring")
}
