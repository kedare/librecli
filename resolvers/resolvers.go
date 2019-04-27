package resolvers

import (
	"encoding/json"
	"fmt"

	"github.com/kedare/librecli/client"
	"github.com/kedare/librecli/entities"
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

		res, err := req.Send()
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
	base := client.BuildAPIClient()
	req := base.Request()
	req.Path(fmt.Sprintf("/api/v0/devices/%v", id))
	res, err := req.Send()
	if err != nil {
		fmt.Println(err)
	}

	getDevicesResponse := entities.GetDevicesResponse{}
	err = json.Unmarshal([]byte(res.String()), &getDevicesResponse)
	return getDevicesResponse.Devices[0]
}
