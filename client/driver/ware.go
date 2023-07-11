package driver

import (
	"context"
	"github.com/go-ole/go-ole/oleutil"
	"ware-data-api/user"
)

func (c *Client) AddWareData(ctx context.Context, ldIndex int, wares []user.WareData) error {
	_, err := oleutil.PutProperty(c.drv, "LDIndex", ldIndex)
	if err != nil {
		return err
	}

	_, err = oleutil.CallMethod(c.drv, "EnumLD")
	if err != nil {
		return err
	}

	_, err = oleutil.CallMethod(c.drv, "SetActiveLD")
	if err != nil {
		return err
	}

	_, err = oleutil.CallMethod(c.drv, "Connect")
	if err != nil {
		return err
	}

	for _, _ = range wares {
		_, err = oleutil.CallMethod(c.drv, "GetEmptyPLUNumber")
		if err != nil {
			return err
		}

	}

	_, err = oleutil.CallMethod(c.drv, "Disconnect")
	if err != nil {
		return err
	}

	return nil
}
