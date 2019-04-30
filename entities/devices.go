package entities

type Device struct {
	ID       int    `json:"device_id"`
	Hostname string `json:"hostname"`
}

type GetDevicesResponse struct {
	Status  string   `json:"status"`
	Devices []Device `json:"devices"`
}
