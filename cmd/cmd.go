package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// OutputFormat defines output format to be used to display the data
var OutputFormat string

// Verbose is the toggle to enable verbose logging
var Verbose bool

var cmdRoot = &cobra.Command{Use: "librecli"}

var cmdBGP = &cobra.Command{
	Use: "bgp",
}

var cmdBGPPeers = &cobra.Command{
	Use: "peers",
}

var cmdBGPPeersList = &cobra.Command{
	Use:    "list <peer asn or device name>",
	Short:  "List BGP peers",
	PreRun: setDebug,
	Run:    ListBGPPeers,
}

var cmdBGPPeersCounters = &cobra.Command{
	Use:    "counters <device name>",
	Short:  "List BGP counters",
	PreRun: setDebug,
	Run:    ListBGPCounters,
}

var cmdFDB = &cobra.Command{
	Use: "fdb",
}

var cmdFDBLookup = &cobra.Command{
	Use:    "lookup [MAC address]",
	Short:  "Look for a MAC in the centralized FDB table",
	PreRun: setDebug,
	Run:    LookupFDB,
	Args:   cobra.MinimumNArgs(1),
}

var cmdIPv4 = &cobra.Command{
	Use: "ipv4",
}

var cmdIPv4Lookup = &cobra.Command{
	Use:    "lookup [IPv4 address]",
	Short:  "Lookup for an IPv4 address on the centralized IP table",
	PreRun: setDebug,
	Run:    LookupIPv4,
	Args:   cobra.MinimumNArgs(1),
}

func setDebug(cmd *cobra.Command, args []string) {
	if Verbose {
		log.SetLevel(log.DebugLevel)
	}
}

// Setup is the function that will register all the CLI handlers and parameters
func Setup(version string) {
	cmdRoot.Version = version

	// Bind Commands
	cmdRoot.AddCommand(cmdBGP)
	cmdBGP.AddCommand(cmdBGPPeers)
	cmdBGPPeers.AddCommand(cmdBGPPeersList)
	cmdBGPPeers.AddCommand(cmdBGPPeersCounters)

	cmdRoot.AddCommand(cmdFDB)
	cmdFDB.AddCommand(cmdFDBLookup)

	cmdRoot.AddCommand(cmdIPv4)
	cmdIPv4.AddCommand(cmdIPv4Lookup)

	cmdRoot.PersistentFlags().StringVarP(&OutputFormat, "format", "f", "table", "Output format: table|list|json")
	cmdRoot.PersistentFlags().BoolVarP(&Verbose, "verbose", "V", false, "Enable verbose mode")

	cmdRoot.Execute()
}
