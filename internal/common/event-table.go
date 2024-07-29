package common

import (
	"github.com/scylladb/gocqlx/v3/table"
)

func NewEventTable() *table.Table {
	m := table.Metadata{
		Name: "events",
		Columns: []string{
			"id",
			"browser_fingerprint",
			"canvas_fingerprint",
			"created_at",
			"device_language",
			"device_timezone",
			"event_name",
			"font_fingerprint",
			"incognito",
			"ip",
			"periodic_wave",
			"processed",
			"screen_resolution",
			"session",
			"status",
			"storage",
			"updated_at",
			"user_agent",
			"user_id",
			"utm",
			"webgl_fingerprint",
		},
		PartKey: []string{"id"},
		SortKey: []string{"created_at"},
	}
	return table.New(m)
}
