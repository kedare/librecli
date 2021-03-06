package resolvers

import (
	"encoding/json"
	"fmt"

	"github.com/kedare/librecli/entities"
	"github.com/kedare/librecli/network"
)

// GetDeviceByID will return the device that matches the given ID
func GetDeviceByID(id int) entities.Device {
	base := network.BuildAPIClient()
	req := base.Request()
	req.Path(fmt.Sprintf("/api/v0/devices/%v", id))
	res, err := network.RunRequestIfNotCached(fmt.Sprintf("nms.devices:%v", id), req)
	if err != nil {
		fmt.Println(err)
	}

	getDevicesResponse := entities.GetDevicesResponse{}
	err = json.Unmarshal([]byte(res.String()), &getDevicesResponse)
	return getDevicesResponse.Devices[0]
}

// GetDevicesByParam will return all the devices that match the specified parameters
func GetDevicesByParam(key string, value string) []entities.Device {
	base := network.BuildAPIClient()
	req := base.Request()
	req.Path("/api/v0/devices")
	req.AddQuery("type", key)
	req.AddQuery("query", value)
	res, err := network.RunRequestIfNotCached(fmt.Sprintf("nms.devices:%v=%v", key, value), req)
	if err != nil {
		fmt.Println(err)
	}

	getDevicesResponse := entities.GetDevicesResponse{}
	err = json.Unmarshal([]byte(res.String()), &getDevicesResponse)
	return getDevicesResponse.Devices
}

// GetIPv4InterfacesByDevices will get the data of the interfaces associated to a specified device
func GetIPv4InterfacesByDevice(hostname string) []entities.IPv4Address {
	base := network.BuildAPIClient()
	req := base.Request()
	req.Path(fmt.Sprintf("/api/v0/devices/%v/ip", hostname))
	res, err := network.RunRequestIfNotCached(fmt.Sprintf("nms.device_ifip:%v", hostname), req)
	if err != nil {
		fmt.Println(err)
	}

	getDeviceIPv4AddressesResponse := entities.GetDeviceIPv4AddressesResponse{}
	err = json.Unmarshal([]byte(res.String()), &getDeviceIPv4AddressesResponse)
	return getDeviceIPv4AddressesResponse.Addresses
}

// GetPortByID will get data for a specified port depending of the given ID
func GetPortByID(id int) entities.Port {
	base := network.BuildAPIClient()
	req := base.Request()
	req.Path(fmt.Sprintf("/api/v0/ports/%v", id))
	res, err := network.RunRequestIfNotCached(fmt.Sprintf("nms.ports:%v", id), req)
	if err != nil {
		fmt.Println(err)
	}

	getPortsResponse := entities.GetPortsResponse{}
	err = json.Unmarshal([]byte(res.String()), &getPortsResponse)
	return getPortsResponse.Ports[0]
}
