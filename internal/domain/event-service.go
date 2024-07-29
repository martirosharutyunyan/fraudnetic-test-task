package domain

import (
	"context"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/dto"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/entity"
)

type EventService interface {
	Insert(ctx context.Context, createEventDto *dto.CreateEvent) (*entity.Event, error)
	Get(ctx context.Context, eventPageDto *dto.EventPageOptions) (*dto.Page[*entity.Event], error)
}
