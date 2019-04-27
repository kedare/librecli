package entities

type ASData struct {
	Holder string `json:"holder"`
}

type ASInfoResponse struct {
	Data ASData `json:"data"`
}

type Device struct {
	ID       int    `json:"device_id"`
	Hostname string `json:"hostname"`
}

type GetDevicesResponse struct {
	Status  string   `json:"status"`
	Devices []Device `json:"devices"`
}

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

