package cmd

import (
	"fmt"

	"github.com/kedare/librecli/entities"
	"github.com/kedare/librecli/outputs"
	"github.com/kedare/librecli/resolvers"
	"github.com/kedare/librecli/utils"
	"github.com/spf13/cobra"
)

// LookupIPv4 is the command handler that allows you to do a lookup in the centralized IPv4 database
// The parameter to give is which IP you want to lookup
func LookupIPv4(cmd *cobra.Command, args []string) {

	spinner := utils.NewSpinner()
	spinner.Start()

	devices := resolvers.GetDevicesByParam("ipv4", args[0])
	allDevicesInterfaces := []entities.IPv4Address{}

	for _, device := range devices {
		deviceInterfaces := resolvers.GetIPv4InterfacesByDevice(device.Hostname)
		for _, deviceInterface := range deviceInterfaces {
			if deviceInterface.IPv4Address == args[0] {
				deviceInterface.DeviceID = device.ID
				allDevicesInterfaces = append(allDevicesInterfaces, deviceInterface)
			}
		}
	}

	headers := []string{"Device", "Port", "IP Address"}
	var data []map[string]string

	for _, deviceInterface := range allDevicesInterfaces {
		data = append(data, map[string]string{
			"Device":     resolvers.GetDeviceByID(deviceInterface.DeviceID).Hostname,
			"Port":       fmt.Sprint(resolvers.GetPortByID(deviceInterface.PortID).Name),
			"IP Address": fmt.Sprintf("%v/%v", deviceInterface.IPv4Address, deviceInterface.IPv4PrefixLen),
		})
	}

	spinner.Stop()
	outputs.OutputAs(OutputFormat, headers, data)
}
