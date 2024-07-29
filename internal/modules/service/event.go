package service

import (
	"context"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/domain"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/dto"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/entity"
	"github.com/scylladb/gocqlx/v3"
)

type Event struct {
	db        *gocqlx.Session
	eventRepo domain.EventRepo
}

var _ domain.EventService = (*Event)(nil)

func (e *Event) Insert(ctx context.Context, createEventDto *dto.CreateEvent) (*entity.Event, error) {
	return e.eventRepo.Insert(ctx, createEventDto)
}

func (e *Event) Get(ctx context.Context, eventPageDto *dto.EventPageOptions) (*dto.Page[*entity.Event], error) {
	return e.eventRepo.Get(ctx, eventPageDto)
}

func NewEvent(eventRepo domain.EventRepo) *Event {
	return &Event{
		eventRepo: eventRepo,
	}
}
