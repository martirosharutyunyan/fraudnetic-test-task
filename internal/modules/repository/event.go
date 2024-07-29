package repository

import (
	"context"
	"github.com/google/uuid"
	httpErrors "github.com/martirosharutyunyan/fraudnetic-test-task/internal/common/http-errors"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/domain"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/dto"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/entity"
	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/qb"
	"github.com/scylladb/gocqlx/v3/table"
	"time"
)

type Event struct {
	db    *gocqlx.Session
	table *table.Table
}

var _ domain.EventRepo = (*Event)(nil)

func (e *Event) Insert(ctx context.Context, createEventDto *dto.CreateEvent) (*entity.Event, error) {
	event := &entity.Event{
		ID:                 uuid.New().String(),
		BrowserFingerprint: createEventDto.BrowserFingerprint,
		CanvasFingerprint:  createEventDto.CanvasFingerprint,
		CreatedAt:          time.Now(),
		DeviceLanguage:     createEventDto.DeviceLanguage,
		DeviceTimezone:     createEventDto.DeviceTimezone,
		EventName:          createEventDto.EventName,
		FontFingerprint:    createEventDto.FontFingerprint,
		Incognito:          createEventDto.Incognito,
		IP:                 createEventDto.Ip,
		PeriodicWave:       createEventDto.PeriodicWave,
		Processed:          createEventDto.Processed,
		ScreenResolution:   createEventDto.ScreenResolution,
		Session:            createEventDto.Session,
		Status:             createEventDto.Status,
		Storage:            createEventDto.Storage,
		UpdatedAt:          time.Now(),
		UserAgent:          createEventDto.UserAgent,
		UserId:             createEventDto.UserId,
		Utm:                createEventDto.Utm,
		WebglFingerprint:   createEventDto.WebglFingerprint,
	}

	err := e.db.Query(e.table.Insert()).BindStruct(event).WithContext(ctx).Exec()
	if err != nil {
		return nil, httpErrors.NewBadRequestError(err.Error())
	}

	return event, nil
}

func (e *Event) Get(ctx context.Context, eventPageDto *dto.EventPageOptions) (*dto.Page[*entity.Event], error) {
	var events []*entity.Event

	stmt := qb.Select(e.table.Metadata().Name).Columns(e.table.Metadata().Columns...)

	values := qb.M{}

	if eventPageDto.EventName != nil {
		stmt = stmt.Where(qb.Eq("event_name"))
		values["event_name"] = *eventPageDto.EventName
	}
	if eventPageDto.DeviceLanguage != nil {
		stmt = stmt.Where(qb.Eq("device_language"))
		values["device_language"] = *eventPageDto.DeviceLanguage
	}
	if eventPageDto.Incognito != nil {
		stmt = stmt.Where(qb.Eq("incognito"))
		values["incognito"] = *eventPageDto.Incognito
	}
	if eventPageDto.IP != nil {
		stmt = stmt.Where(qb.Eq("ip"))
		values["ip"] = *eventPageDto.IP
	}
	if eventPageDto.EventType != nil {
		stmt = stmt.Where(qb.Eq("event_type"))
		values["event_type"] = *eventPageDto.EventType
	}

	if eventPageDto.Order != nil {
		var order qb.Order
		if *eventPageDto.Order == "ASC" {
			order = qb.ASC
		} else {
			order = qb.DESC
		}
		stmt = stmt.OrderBy("created_at", order)
	}

	var count int
	countStmt, names := qb.Select("events").CountAll().ToCql()
	err := e.db.ContextQuery(ctx, countStmt, names).Scan(&count)
	if err != nil {
		return nil, httpErrors.NewBadRequestError(err.Error())
	}

	stmt = stmt.Limit(eventPageDto.Limit).AllowFiltering()

	err = e.db.Query(stmt.ToCql()).BindMap(values).SelectRelease(&events)
	if err != nil {
		return nil, httpErrors.NewBadRequestError(err.Error())
	}

	return dto.NewPage(events, count), nil
}

func NewEvent(db *gocqlx.Session, table *table.Table) *Event {
	return &Event{
		db:    db,
		table: table,
	}
}
