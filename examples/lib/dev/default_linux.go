package dev

import (
	"github.com/k-katsuda/ble"
	"github.com/k-katsuda/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
