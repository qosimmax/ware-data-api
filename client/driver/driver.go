package driver

import (
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"ware-data-api/config"
)

type Client struct {
	drv *ole.IDispatch
}

func (c *Client) Init(config *config.Config) error {
	err := ole.CoInitialize(0)
	if err != nil {
		return err
	}

	iUnknown, err := oleutil.CreateObject("AddIn.DrvLP")
	if err != nil {
		return err
	}

	c.drv, err = iUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Close() error {
	c.drv.Release()
	ole.CoUninitialize()
	return nil
}
