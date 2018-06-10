package api

// AppInfo is shared wallpaperize info type
type AppInfo struct {
	AppVersion string `json:"app_version"`
	Arch       string `json:"arch"`
	OS         string `json:"os"`
}
