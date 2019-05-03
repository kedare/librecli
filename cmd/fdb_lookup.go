package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/kedare/librecli/entities"
	"github.com/kedare/librecli/network"
	"github.com/kedare/librecli/outputs"
	"github.com/kedare/librecli/resolvers"
	"github.com/spf13/cobra"
)

// LookupFDB is the command handler that will allow you to do a lookup in the centralized FDB table in LibreNMS
// The paramter to give it is the MAC address (Without any special character) you want to lookup
func LookupFDB(cmd *cobra.Command, args []string) {
	base := network.BuildAPIClient()
	req := base.Request()
	req.Path(fmt.Sprintf("/api/v0/resources/fdb/%v", args[0]))

	res, err := req.Send()
	if err != nil {
		fmt.Println(err)
	}

	fdbLookupResponse := entities.FDBLookupResponse{}
	err = json.Unmarshal([]byte(res.String()), &fdbLookupResponse)

	headers := []string{"Device", "Port", "VLAN"}
	var data []map[string]string

	for _, fdbPort := range fdbLookupResponse.FDBPorts {
		data = append(data, map[string]string{
			"Device": resolvers.GetDeviceByID(fdbPort.DeviceID).Hostname,
			"Port":   resolvers.GetPortByID(fdbPort.PortID).Name,
			"VLAN":   fmt.Sprint(fdbPort.VlanID),
		})
	}
	outputs.OutputAs(OutputFormat, headers, data)
}
