package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/kedare/librecli/colorizers"
	"github.com/kedare/librecli/entities"
	"github.com/kedare/librecli/network"
	"github.com/kedare/librecli/outputs"
	"github.com/kedare/librecli/resolvers"
	"github.com/spf13/cobra"
)

func ListBGPCounters(cmd *cobra.Command, args []string) {
	base := network.BuildAPIClient()
	req := base.Request()
	req.Path("/api/v0/routing/bgp/cbgp")

	if len(args) > 0 {
		req.AddQuery("hostname", args[0])
	}

	res, err := req.Send()
	if err != nil {
		fmt.Println(err)
	}

	listBgpCountersResponse := entities.ListBGPCountersResponse{}
	err = json.Unmarshal([]byte(res.String()), &listBgpCountersResponse)

	var data []map[string]string

	for _, bgpCounter := range listBgpCountersResponse.BGPCounters {
		data = append(data,
			map[string]string{
				"Device":       resolvers.GetDeviceByID(bgpCounter.DeviceID).Hostname,
				"Peer IP":      bgpCounter.BGPPeerIdentifier,
				"Accepted Pfx": colorizers.ShouldBeHigherThan(bgpCounter.AcceptedPrefixes, 0),
				"Denied Pfx":   colorizers.ShouldBeLowerThan(bgpCounter.DeniedPrefixes, 1),
				"Pfx Thrsd":    fmt.Sprint(bgpCounter.PrefixThreshold),
				"Adv Pfx":      colorizers.ShouldBeHigherThan(bgpCounter.AdvertisedPrefixes, 0),
			})
	}
	outputs.OutputAs(OutputFormat, data)
}
