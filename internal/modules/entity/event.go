package entity

import "time"

type Event struct {
	ID                 string    `db:"id"`
	BrowserFingerprint int       `db:"browser_fingerprint"`
	CanvasFingerprint  int       `db:"canvas_fingerprint"`
	CreatedAt          time.Time `db:"created_at"`
	DeviceLanguage     string    `db:"device_language"`
	DeviceTimezone     int       `db:"device_timezone"`
	EventName          string    `db:"event_name"`
	FontFingerprint    int       `db:"font_fingerprint"`
	Incognito          bool      `db:"incognito"`
	IP                 string    `db:"ip"`
	PeriodicWave       int       `db:"periodic_wave"`
	Processed          bool      `db:"processed"`
	ScreenResolution   string    `db:"screen_resolution"`
	Session            string    `db:"session"`
	Status             bool      `db:"status"`
	Storage            int       `db:"storage"`
	UpdatedAt          time.Time `db:"updated_at"`
	UserAgent          string    `db:"user_agent"`
	UserId             int       `db:"user_id"`
	Utm                string    `db:"utm"`
	WebglFingerprint   int       `db:"webgl_fingerprint"`
}
