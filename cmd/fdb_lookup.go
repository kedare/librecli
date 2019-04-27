package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/apcera/termtables"
	"github.com/kedare/librecli/client"
	"github.com/kedare/librecli/resolvers"
	"github.com/kedare/librecli/entities"
	"github.com/spf13/cobra"
)



func LookupFDB(cmd *cobra.Command, args []string) {
	base := client.BuildAPIClient()
	req := base.Request()
	req.Path(fmt.Sprintf("/api/v0/resources/fdb/%v", args[0]))

	res, err := req.Send()
	if err != nil {
		fmt.Println(err)
	}

	fdbLookupResponse := entities.FDBLookupResponse{}
	err = json.Unmarshal([]byte(res.String()), &fdbLookupResponse)
	table := termtables.CreateTable()
	table.AddHeaders("Device", "Port", "VLAN")

	for _, fdbPort := range fdbLookupResponse.FDBPorts {
		table.AddRow(
			resolvers.GetDeviceByID(fdbPort.DeviceID).Hostname,
			fdbPort.PortID,
			fdbPort.VlanID)
	}
	fmt.Println(table.Render())
}
