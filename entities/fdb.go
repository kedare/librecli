package entities

// FDBPort represent a port in the centralizd FDB table
type FDBPort struct {
	FDBPortID  int    `json:"ports_fdb_id"`
	PortID     int    `json:"port_id"`
	MacAddress string `json:"mac_address"`
	VlanID     int    `json:"vlan_id"`
	DeviceID   int    `json:"device_id"`
}

// FDBLookupResponse represents a reponse from /api/v0/resources/fdb
type FDBLookupResponse struct {
	Status   string    `json:"status"`
	FDBPorts []FDBPort `json:"ports_fdb"`
	Count    int       `json:"count"`
}
