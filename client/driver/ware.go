package driver

import (
	"context"
	"fmt"
	"ware-data-api/user"

	"github.com/go-ole/go-ole/oleutil"
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

	code, err := oleutil.CallMethod(c.drv, "Connect")
	if err != nil {
		return err
	}

	if code.Value().(int32) != 0 {
		return fmt.Errorf("not connected")
	}

	for _, ware := range wares {
		_, err = oleutil.PutProperty(c.drv, "PLUNumber", ware.ItemCode)
		if err != nil {
			return err
		}

		_, err = oleutil.PutProperty(c.drv, "Price", ware.Price)
		if err != nil {
			return err
		}

		_, err = oleutil.PutProperty(c.drv, "ItemCode", ware.ItemCode)
		if err != nil {
			return err
		}

		_, err = oleutil.PutProperty(c.drv, "NameFirst", ware.Name)
		if err != nil {
			return err
		}

		_, err = oleutil.PutProperty(c.drv, "ShelfLife", 0)
		if err != nil {
			return err
		}

		_, err = oleutil.PutProperty(c.drv, "GroupCode", 0)
		if err != nil {
			return err
		}

		_, err = oleutil.PutProperty(c.drv, "GoodsType", ware.GoodsType)
		if err != nil {
			return err
		}

		_, err = oleutil.PutProperty(c.drv, "BCFormat", 7)
		if err != nil {
			return err
		}

		_, err = oleutil.CallMethod(c.drv, "SetPLUDataEx")
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
