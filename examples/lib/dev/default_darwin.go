package dev

import (
	"github.com/k-katsuda/ble"
	"github.com/k-katsuda/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
