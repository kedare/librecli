package entities

// BGPSession represents a BGP session as it is returned by the LibreNMS API
type BGPSession struct {
	BGPPeerID                  int    `json:"bgpPeer_id"`
	DeviceID                   int    `json:"device_id"`
	ASText                     string `json:"astext"`
	BGPPeerIdentifier          string `json:"bgpPeerIdentifier"`
	BGPPeerRemoteAS            int    `json:"bgpPeerRemoteAs"`
	BGPPeerState               string `json:"bgpPeerState"`
	BGPPeerAdminStatus         string `json:"bgpPeerAdminStatus"`
	BGPLocalAddr               string `json:"bgpLocalAddr"`
	BGPPeerRemoteAddr          string `json:"bgpPeerRemoteAddr"`
	BGPPeerInUpdates           int    `json:"bgpPeerInUpdates"`
	BGPPeerOutUpdates          int    `json:"bgpPeerOutUpdates"`
	BGPPeerInTotalMessages     int    `json:"bgpPeerInTotalMessages"`
	BGPPeerOutTotalMessages    int    `json:"bgpPeerOutTotalMessages"`
	BGPPeerFMSEstablishedTime  int    `json:"bgpPeerFmsEstablishedTime"`
	BGPPeerInUpdateElapsedTime int    `json:"bgpPeerInUpdateElapsedTime"`
	ContextName                string `json:"contextName"`
}

// ListBGPResponse represent an API response call to /api/v0/bgp
type ListBGPResponse struct {
	Status      string       `json:"status"`
	Message     string       `json:"message"`
	BGPSessions []BGPSession `json:"bgp_sessions"`
	Count       int          `json:"count"`
}

// BGPCounters represent BGP counters as they are returned by the LibreNMS API
type BGPCounters struct {
	DeviceID             int    `json:"device_id"`
	BGPPeerIdentifier    string `json:"bgpPeerIdentifier"`
	AFI                  string `json:"afi"`
	SAFI                 string `json:"safi"`
	AcceptedPrefixes     int    `json:"AcceptedPrefixes"`
	DeniedPrefixes       int    `json:"DeniedPrefixes"`
	PrefixAdminLimit     int    `json:"PrefixAdminLimit"`
	PrefixThreshold      int    `json:"PrefixThreshold"`
	PrefixClearThreshold int    `json:"PrefixClearThreshold"`
	AdvertisedPrefixes   int    `json:"AdvertisedPrefixes"`
	SuppressedPrefixes   int    `json:"SupressedPrefixes"`
	WithdrawnPrefixes    int    `json:"WithdrawnPrefixes"`
}

// ListBGPCountersResponse represent an API response call to /api/v0/routing/bgp/cbgp
type ListBGPCountersResponse struct {
	Status      string        `json:"status"`
	BGPCounters []BGPCounters `json:"bgp_counters"`
	Count       int           `json:"count"`
}

// ASData represent the data associated with an AS
type ASData struct {
	Holder string `json:"holder"`
}

// ASInfoResponse represent the response to a query to https://stats.ripe.net/data/as-overview/data.json
type ASInfoResponse struct {
	Data ASData `json:"data"`
}
