package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/h2non/gentleman.v2"

	"github.com/kedare/librecli/colorizers"
	"github.com/kedare/librecli/entities"
	"github.com/kedare/librecli/network"
	"github.com/kedare/librecli/outputs"
	"github.com/kedare/librecli/resolvers"
	"github.com/spf13/cobra"
)

// ListBGPPeers is the command handler that will display the configured BGP peers/session on all the devices by default.
// If a parameter is given, it can either be an ASN to filter by remote ASN or a device hostname to filter by device
func ListBGPPeers(cmd *cobra.Command, args []string) {
	base := network.BuildAPIClient()
	req := base.Request()
	req.Path("/api/v0/bgp")

	var res *gentleman.Response
	var err error

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
		res, err = network.RunRequestIfNotCached(fmt.Sprintf("nms.bgp:%v", args[0]), req)
	} else {
		res, err = network.RunRequestIfNotCached("nms.bgp", req)
	}

	if err != nil {
		fmt.Println(err)
	}

	listBgpResponse := entities.ListBGPResponse{}
	err = json.Unmarshal([]byte(res.String()), &listBgpResponse)

	headers := []string{"Device", "Local IP", "Peer IP", "Peer AS", "AS Holder", "State", "Admin State"}
	var data []map[string]string

	for _, bgpSession := range listBgpResponse.BGPSessions {
		data = append(data,
			map[string]string{
				"Device":      resolvers.GetDeviceByID(bgpSession.DeviceID).Hostname,
				"Local IP":    bgpSession.BGPLocalAddr,
				"Peer IP":     bgpSession.BGPPeerIdentifier,
				"Peer AS":     fmt.Sprint(bgpSession.BGPPeerRemoteAS),
				"AS Holder":   resolvers.GetASHolderByASN(bgpSession.BGPPeerRemoteAS),
				"State":       colorizers.ColorizeBGPPeerStatus(bgpSession.BGPPeerState),
				"Admin State": colorizers.ColorizeBGPPeerAdminStatus(bgpSession.BGPPeerAdminStatus)},
		)
	}
	outputs.OutputAs(OutputFormat, headers, data)
}
