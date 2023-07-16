package user

import "context"

type WareData struct {
	ItemCode  int
	Name      string
	Price     int
	Count     int
	GoodsType int
}

// WareDataAdder  is an interface for adding ware data to device
type WareDataAdder interface {
	AddWareData(ctx context.Context, ldIndex int, wares []WareData) error
}

type DeviceWareAdder interface {
	LDDeviceGetter
	WareDataAdder
}
