package dto

type CreateEvent struct {
	BrowserFingerprint int    `json:"browser_fingerprint"`
	CanvasFingerprint  int    `json:"canvas_fingerprint"`
	CreatedAt          string `json:"created_at"`
	DeviceLanguage     string `json:"device_language"`
	DeviceTimezone     int    `json:"device_timezone"`
	EventName          string `json:"event_name"`
	FontFingerprint    int    `json:"font_fingerprint"`
	Incognito          bool   `json:"incognito"`
	Ip                 string `json:"ip"`
	PeriodicWave       int    `json:"periodic_wave"`
	Processed          bool   `json:"processed"`
	ScreenResolution   string `json:"screen_resolution"`
	Session            string `json:"session"`
	Status             bool   `json:"status"`
	Storage            int    `json:"storage"`
	UpdatedAt          string `json:"updated_at"`
	UserAgent          string `json:"user_agent"`
	UserId             int    `json:"user_id"`
	Utm                string `json:"utm"`
	WebglFingerprint   int    `json:"webgl_fingerprint"`
}
