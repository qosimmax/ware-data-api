package user

import "context"

type WareData struct {
	Code  int
	Price float32
	Count int
	Type  int
}

// WareDataAdder  is an interface for adding ware data to device
type WareDataAdder interface {
	AddWareData(ctx context.Context, ldNumber int, wares []WareData) error
}
