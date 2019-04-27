package main

import (
	"github.com/kedare/librecli/cmd"
	"github.com/spf13/cobra"
)

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
	Run:   cmd.ListBGPPeers,
}

var cmdBGPPeersCounters = &cobra.Command{
	Use:   "counters <device name>",
	Short: "List BGP counters",
	Run:   cmd.ListBGPCounters,
}

var cmdFDB = &cobra.Command{
	Use: "fdb",
}

var cmdFDBLookup = &cobra.Command{
	Use:   "lookup [MAC address]",
	Short: "Look for a MAC in the centralized FDB table",
	Run:   cmd.LookupFDB,
	Args:  cobra.MinimumNArgs(1),
}

func main() {

	// Bind Commands
	cmdRoot.AddCommand(cmdBGP)
	cmdBGP.AddCommand(cmdBGPPeers)
	cmdBGPPeers.AddCommand(cmdBGPPeersList)
	cmdBGPPeers.AddCommand(cmdBGPPeersCounters)

	cmdRoot.AddCommand(cmdFDB)
	cmdFDB.AddCommand(cmdFDBLookup)

	cmdRoot.Execute()
}
