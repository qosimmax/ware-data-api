package driver

import (
	"ware-data-api/config"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type Client struct {
	drv *ole.IDispatch
}

func (c *Client) Init(config *config.Config) error {
	err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
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
