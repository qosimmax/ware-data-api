package user

import "context"

type WareData struct {
	Code  int
	Price float32
	Count int
	Type  int
}

type WareDataAdder interface {
	AddWareData(ctx context.Context, wares []WareData) error
}
