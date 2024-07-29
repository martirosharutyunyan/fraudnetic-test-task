package dto

type EventPageOptions struct {
	DeviceLanguage *string `query:"device_language"`
	EventName      *string `query:"event_name"`
	Incognito      *bool   `query:"incognito"`
	IP             *string `query:"ip"`
	EventType      *string `query:"event_type"`
	Order          *string `query:"order"`
	Limit          uint    `query:"limit"`
	Offset         uint    `query:"offset"`
}
