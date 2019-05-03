package entities

// Port represent a physical (or logical) port
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

// GetPortResponse represents a response from /api/v0/port
type GetPortsResponse struct {
	Status string `json:"status"`
	Ports  []Port `json:"port"`
}

// IPv4Address represent an IPv4 Address
type IPv4Address struct {
	IPv4AddressID int    `json:"ipv4_address_id"`
	IPv4Address   string `json:"ipv4_address"`
	IPv4PrefixLen int    `json:"ipv4_prefixlen"`
	IPv4NetworkID int    `json:"ipv4_network_id"`
	PortID        int    `json:"port_id"`
	DeviceID      int    `json:"device_id"`
	ContextName   string `json:"context_name"`
}

// GetDeviceIPv4AddressResponse represents a response from /api/v0/devices/:hostname:/ip
type GetDeviceIPv4AddressesResponse struct {
	Status    string        `json:"status"`
	Message   string        `json:"message"`
	Addresses []IPv4Address `json:"addresses"`
}
