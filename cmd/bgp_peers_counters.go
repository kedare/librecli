package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kedare/librecli/colorizers"
	"github.com/kedare/librecli/entities"
	"github.com/kedare/librecli/network"
	"github.com/kedare/librecli/outputs"
	"github.com/kedare/librecli/resolvers"
	"github.com/kedare/librecli/utils"
	"github.com/spf13/cobra"
	"gopkg.in/h2non/gentleman.v2"
)

// ListBGPCounters is the command handler that will display the counters coming from the BGP sessions.
// A hostname can be given as parameter to filter by device
func ListBGPCounters(cmd *cobra.Command, args []string) {
	base := network.BuildAPIClient()
	req := base.Request()
	req.Path("/api/v0/routing/bgp/cbgp")

	var res *gentleman.Response
	var err error

	spinner := utils.NewSpinner()
	spinner.Start()

	if len(args) > 0 {
		// Try to get arg as ASN
		asn, err := strconv.Atoi(args[0])
		if err != nil {
			// Failed to convert to INT, this is not an ASN, so this must be a device name
			req.AddQuery("hostname", args[0])
		} else {
			// Conversion is good, this is an ASN
			req.AddQuery("remote_asn", fmt.Sprint(asn))
		}
		res, err = network.RunRequestIfNotCached(fmt.Sprintf("nms.cbgp:%v", args[0]), req)
	} else {
		res, err = network.RunRequestIfNotCached("nms.cbgp", req)
	}

	if err != nil {
		fmt.Println(err)
		spinner.Stop()
	}

	listBgpCountersResponse := entities.ListBGPCountersResponse{}
	err = json.Unmarshal([]byte(res.String()), &listBgpCountersResponse)

	headers := []string{"Device", "Peer IP", "Accepted Pfx", "Denied Pfx", "Pfx Thrsd", "Adv Pfx"}
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

	spinner.Stop()
	outputs.OutputAs(OutputFormat, headers, data)
}
