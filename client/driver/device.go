package driver

import (
	"context"
	"net"
	"ware-data-api/user"

	"github.com/go-ole/go-ole/oleutil"
)

func (c *Client) FindActiveDevices(ctx context.Context, fromIP, toIP string) ([]user.DeviceData, error) {
	startIP := net.ParseIP(fromIP)
	endIP := net.ParseIP(toIP)
	var (
		devices []user.DeviceData
	)

	for ip := startIP; !ip.To4().Equal(endIP.To4()); inc(ip) {
		_, err := oleutil.PutProperty(c.drv, "RemoteHost", ip.String())
		if err != nil {
			return nil, err
		}

		_, err = oleutil.PutProperty(c.drv, "RemotePort", 1111)
		if err != nil {
			return nil, err
		}

		_, err = oleutil.PutProperty(c.drv, "TimeoutUDP", 500)
		if err != nil {
			return nil, err
		}

		_, err = oleutil.PutProperty(c.drv, "DeviceInterface", 1)
		if err != nil {
			return nil, err
		}

		_, err = oleutil.CallMethod(c.drv, "Connect")
		if err != nil {
			return nil, err
		}

		val, err := oleutil.GetProperty(c.drv, "Connected")
		if err != nil {
			return nil, err
		}

		_, err = oleutil.CallMethod(c.drv, "Beep")
		if err != nil {
			return nil, err
		}

		_, err = oleutil.CallMethod(c.drv, "Disconnect")
		if err != nil {
			return nil, err
		}

		if val == nil {
			continue
		}

		if val.Value().(bool) {
			devices = append(devices, user.DeviceData{
				Number:     0,
				Name:       "",
				RemoteHost: ip.String(),
				RemotePort: 1111,
			})
		}

	}

	return devices, nil
}

func (c *Client) GetCountLDDevice(ctx context.Context) (int, error) {
	_, err := oleutil.CallMethod(c.drv, "GetCountLD")
	if err != nil {
		return 0, err
	}

	val, err := oleutil.GetProperty(c.drv, "LDCount")
	if err != nil {
		return 0, err
	}

	if val == nil {
		return 0, nil
	}

	return int(val.Val), nil

}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
