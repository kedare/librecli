package entities


type FDBPort struct {
	FDBPortID  int    `json:"ports_fdb_id"`
	PortID     int    `json:"port_id"`
	MacAddress string `json:"mac_address"`
	VlanID     int    `json:"vlan_id"`
	DeviceID   int    `json:"device_id"`
}

type FDBLookupResponse struct {
	Status   string    `json:"status"`
	FDBPorts []FDBPort `json:"ports_fdb"`
	Count    int       `json:"count"`
}