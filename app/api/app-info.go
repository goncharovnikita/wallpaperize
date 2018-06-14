package api

// AppInfo is shared wallpaperize info type
type AppInfo struct {
	AppVersion   string   `json:"app_version"`
	Arch         string   `json:"arch"`
	OS           string   `json:"os"`
	Build        string   `json:"build"`
	DailyImages  []string `json:"daily_images"`
	RandomImages []string `json:"random_images"`
}
