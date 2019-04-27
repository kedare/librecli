package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/apcera/termtables"
	"github.com/kedare/librecli/client"
	"github.com/kedare/librecli/colorizers"
	"github.com/kedare/librecli/entities"
	"github.com/kedare/librecli/resolvers"
	"github.com/spf13/cobra"
)

func ListBGPCounters(cmd *cobra.Command, args []string) {
	base := client.BuildAPIClient()
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
	table := termtables.CreateTable()
	table.AddHeaders("Device", "Peer IP", "Accepted Pfx", "Denied Pfx", "Pfx Thrsd", "Adv Pfx")

	for _, bgpCounter := range listBgpCountersResponse.BGPCounters {
		table.AddRow(
			resolvers.GetDeviceByID(bgpCounter.DeviceID).Hostname,
			bgpCounter.BGPPeerIdentifier,
			colorizers.ShouldBeHigherThan(bgpCounter.AcceptedPrefixes, 0),
			colorizers.ShouldBeLowerThan(bgpCounter.DeniedPrefixes, 1),
			bgpCounter.PrefixThreshold,
			colorizers.ShouldBeHigherThan(bgpCounter.AdvertisedPrefixes, 0))
	}
	fmt.Println(table.Render())
}
