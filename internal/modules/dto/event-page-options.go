package dto

type EventPageOptions struct {
	DeviceLanguage *string `json:"device_language"`
	EventName      *string `json:"event_name"`
	Incognito      *bool   `json:"incognito"`
	IP             *string `json:"ip"`
	EventType      *string `json:"event_type"`
	Order          *string `json:"order"`
	Limit          uint    `json:"limit"`
	Offset         uint    `json:"offset"`
}
