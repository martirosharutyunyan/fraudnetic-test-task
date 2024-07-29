package dto

import (
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/entity"
	"time"
)

type Event struct {
	Id                 string    `json:"id"`
	BrowserFingerprint int       `json:"browserFingerprint"`
	CanvasFingerprint  int       `json:"canvasFingerprint"`
	CreatedAt          time.Time `json:"createdAt"`
	DeviceLanguage     string    `json:"deviceLanguage"`
	DeviceTimezone     int       `json:"deviceTimezone"`
	EventName          string    `json:"eventName"`
	FontFingerprint    int       `json:"fontFingerprint"`
	Incognito          bool      `json:"incognito"`
	IP                 string    `json:"ip"`
	PeriodicWave       int       `json:"periodicWave"`
	Processed          bool      `json:"processed"`
	ScreenResolution   string    `json:"screenResolution"`
	Session            string    `json:"session"`
	Status             bool      `json:"status"`
	Storage            int       `json:"storage"`
	UpdatedAt          time.Time `json:"updatedAt"`
	UserAgent          string    `json:"userAgent"`
	UserId             int       `json:"userId"`
	Utm                string    `json:"utm"`
	WebglFingerprint   int       `json:"webglFingerprint"`
}

func NewEvent(entity *entity.Event) *Event {
	return &Event{
		Id:                 entity.ID,
		BrowserFingerprint: entity.BrowserFingerprint,
		CanvasFingerprint:  entity.CanvasFingerprint,
		CreatedAt:          entity.CreatedAt,
		DeviceLanguage:     entity.DeviceLanguage,
		DeviceTimezone:     entity.DeviceTimezone,
		EventName:          entity.EventName,
		FontFingerprint:    entity.FontFingerprint,
		Incognito:          entity.Incognito,
		IP:                 entity.IP,
		PeriodicWave:       entity.PeriodicWave,
		Processed:          entity.Processed,
		ScreenResolution:   entity.ScreenResolution,
		Session:            entity.Session,
		Status:             entity.Status,
		Storage:            entity.Storage,
		UpdatedAt:          entity.UpdatedAt,
		UserAgent:          entity.UserAgent,
		UserId:             entity.UserId,
		Utm:                entity.Utm,
		WebglFingerprint:   entity.WebglFingerprint,
	}
}
