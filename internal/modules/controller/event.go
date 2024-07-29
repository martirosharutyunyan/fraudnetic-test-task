package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/domain"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/dto"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/shared/request"
	validate "github.com/martirosharutyunyan/fraudnetic-test-task/internal/validation"
)

type Event struct {
	eventService domain.EventService
}

func NewEvent(eventService domain.EventService) *Event {
	return &Event{
		eventService: eventService,
	}
}

func (c *Event) Register(router fiber.Router) {
	event := router.Group("/events")

	event.Post("/", request.Handler(validate.Body(c.PushEvent)))
	event.Get("/", request.Handler(validate.Query(c.GetAll)))
}

func (c *Event) PushEvent(ctx *request.Context, createEventDto *dto.CreateEvent) {
	event, err := c.eventService.Insert(ctx, createEventDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.OK(event)
}

func (c *Event) GetAll(ctx *request.Context, options *dto.EventPageOptions) {
	eventPage, err := c.eventService.Get(ctx, options)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.OK(eventPage)
}
