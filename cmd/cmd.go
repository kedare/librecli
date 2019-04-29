package cmd

import (
	"github.com/spf13/cobra"
)

var OutputFormat string

var cmdRoot = &cobra.Command{Use: "librecli"}

var cmdBGP = &cobra.Command{
	Use: "bgp",
}

var cmdBGPPeers = &cobra.Command{
	Use: "peers",
}

var cmdBGPPeersList = &cobra.Command{
	Use:   "list <peer asn or device name>",
	Short: "List BGP peers",
	Run:   ListBGPPeers,
}

var cmdBGPPeersCounters = &cobra.Command{
	Use:   "counters <device name>",
	Short: "List BGP counters",
	Run:   ListBGPCounters,
}

var cmdFDB = &cobra.Command{
	Use: "fdb",
}

var cmdFDBLookup = &cobra.Command{
	Use:   "lookup [MAC address]",
	Short: "Look for a MAC in the centralized FDB table",
	Run:   LookupFDB,
	Args:  cobra.MinimumNArgs(1),
}

func Setup() {
	// Bind Commands
	cmdRoot.AddCommand(cmdBGP)
	cmdBGP.AddCommand(cmdBGPPeers)
	cmdBGPPeers.AddCommand(cmdBGPPeersList)
	cmdBGPPeers.AddCommand(cmdBGPPeersCounters)

	cmdRoot.AddCommand(cmdFDB)
	cmdFDB.AddCommand(cmdFDBLookup)

	cmdRoot.PersistentFlags().StringVarP(&OutputFormat, "format", "f", "table", "Output format: table|list|json")
	cmdRoot.Execute()
}
