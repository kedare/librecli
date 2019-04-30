package entities

type Port struct {
	ID               int    `json:"port_id"`
	DeviceID         int    `json:"device_id"`
	IfDescr          string `json:"ifDescr"`
	Name             string `json:"ifName"`
	Index            int    `json:"ifIndex"`
	Speed            int64  `json:"ifSpeed"`
	Mtu              int    `json:"ifMtu"`
	Duplex           string `json:"ifDuplex"`
	HighSpeed        int    `json:"ifHighSpeed"`
	ConnectorPresent bool   `json:"ifConnectorPresent"`
	OperStatus       string `json:"ifOperStatus"`
	AdminStatus      string `json:"ifAdminStatus"`
}

type GetPortsResponse struct {
	Status string `json:"status"`
	Ports  []Port `json:"port"`
}
type IPv4Address struct {
	IPv4AddressID int    `json:"ipv4_address_id"`
	IPv4Address   string `json:"ipv4_address"`
	IPv4PrefixLen int    `json:"ipv4_prefixlen"`
	IPv4NetworkID int    `json:"ipv4_network_id"`
	PortID        int    `json:"port_id"`
	DeviceID      int    `json:"device_id"`
	ContextName   string `json:"context_name"`
}

type GetDeviceIPv4AddressesResponse struct {
	Status    string        `json:"status"`
	Message   string        `json:"message"`
	Addresses []IPv4Address `json:"addresses"`
}
