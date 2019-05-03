package resolvers

import (
	"encoding/json"
	"fmt"
	"github.com/kedare/librecli/entities"
	"github.com/kedare/librecli/network"
)

// GetASHolderByASN will run a query to RIPE NCC API to get the Holder name of a specified ASN
func GetASHolderByASN(asn int) string {
	if asn > 64512 && asn < 65535 {
		return "Private ASN"
	} else {
		base := network.BuildHTTPClient()
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
