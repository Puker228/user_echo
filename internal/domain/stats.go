package domain

type UserStats struct {
	AndroidVersion string `json:"android_version"`
	DeviceModel    string `json:"device_model"`
	Manufacturer   string `json:"manufacturer"`
	TotalRamGB     int    `json:"total_ram_gb"`
	AppVersion     string `json:"app_version"`
}
