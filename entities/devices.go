package entities

// Device represent a Device (obviously)
type Device struct {
	ID       int    `json:"device_id"`
	Hostname string `json:"hostname"`
}


// GetDevicesResponse represents the response from LibreNMS to /api/v0/devices
type GetDevicesResponse struct {
	Status  string   `json:"status"`
	Devices []Device `json:"devices"`
}
