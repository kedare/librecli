package cmd

import (
	"encoding/json"
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"strconv"

	"github.com/apcera/termtables"
	"github.com/kedare/librecli/colorizers"
	"github.com/kedare/librecli/entities"
	"github.com/kedare/librecli/network"
	"github.com/kedare/librecli/resolvers"
	"github.com/spf13/cobra"
)

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
			req.AddQuery("asn", fmt.Sprint(asn))
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
	table := termtables.CreateTable()
	table.AddHeaders("Device", "Local IP", "Peer IP", "Peer AS", "AS Holder", "State", "Admin State")

	for _, bgpSession := range listBgpResponse.BGPSessions {
		table.AddRow(
			resolvers.GetDeviceByID(bgpSession.DeviceID).Hostname,
			bgpSession.BGPLocalAddr,
			bgpSession.BGPPeerIdentifier,
			bgpSession.BGPPeerRemoteAS,
			resolvers.GetASHolderByASN(bgpSession.BGPPeerRemoteAS),
			colorizers.ColorizeBGPPeerStatus(bgpSession.BGPPeerState),
			colorizers.ColorizeBGPPeerAdminStatus(bgpSession.BGPPeerAdminStatus))
	}
	fmt.Println(table.Render())
}
