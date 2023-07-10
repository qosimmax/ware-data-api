package user

import "context"

type DeviceData struct {
	Number     int
	Name       string
	RemoteHost string
	RemotePort int
}

// DeviceFinder is an interface for getting active devices
type DeviceFinder interface {
	FindActiveDevices(ctx context.Context) ([]DeviceData, error)
}

// LDDeviceAdder is an interface for adding an LD device
type LDDeviceAdder interface {
	AddLDDevice(ctx context.Context, data DeviceData) error
}

// LDDeviceGetter is an interface for getting a number of logical devices
type LDDeviceGetter interface {
	GetCountLDDevice(ctx context.Context) (int, error)
}
