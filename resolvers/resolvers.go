package resolvers

import (
	"encoding/json"
	"fmt"

	"github.com/kedare/librecli/entities"
	"github.com/kedare/librecli/network"
	"gopkg.in/h2non/gentleman.v2"
)

func GetASHolderByASN(asn int) string {
	if asn > 64512 && asn < 65535 {
		return "Private ASN"
	} else {
		base := gentleman.New()
		base.BaseURL("https://stat.ripe.net/")
		req := base.Request()
		req.Path("/data/as-overview/data.json")
		req.AddQuery("resource", fmt.Sprint(asn))

		res, err := network.RunRequestIfNotCached(fmt.Sprintf("ripe.asn:%v", asn), req)
		if err != nil {
			fmt.Println(err)
		}

		asInfoResponse := entities.ASInfoResponse{}
		err = json.Unmarshal([]byte(res.String()), &asInfoResponse)
		if err != nil {
			fmt.Println(err)
		}
		return asInfoResponse.Data.Holder
	}
}

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
