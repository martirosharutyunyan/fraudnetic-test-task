package request

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	httpErrors "github.com/martirosharutyunyan/fraudnetic-test-task/internal/common/http-errors"
	"net/http"
)

type Context struct {
	*fiber.Ctx
	context.Context
	traceId string
}

func Handler(handler func(ctx *Context)) fiber.Handler {
	return func(context *fiber.Ctx) error {
		ctx := &Context{Context: context.Context(), Ctx: context}

		handler(ctx)
		return nil
	}
}

func (ctx *Context) SetUser(user any) {
	ctx.Locals("user", user)
}

func (ctx *Context) GetUser() any {
	return ctx.Locals("user")
}

func (ctx *Context) StatusOk() {
	ctx.Status(http.StatusOK)
}

func (ctx *Context) Error(err error) {
	var httpError *httpErrors.HTTPError
	if err != nil && errors.As(err, &httpError) {
		ctx.Status(httpError.StatusCode).JSON(httpError)
		return
	}

	ctx.Status(
		http.StatusInternalServerError,
	).JSON(
		httpErrors.NewHTTPError(
			http.StatusInternalServerError,
			"Error is not httpError type please fix",
		),
	)
}

func (ctx *Context) SetTraceId(id string) {
	ctx.traceId = id
}

func (ctx *Context) GetTraceId() string {
	return ctx.traceId
}

func (ctx *Context) OK(body any) {
	ctx.Status(http.StatusOK).JSON(body)
}

func (ctx *Context) BadRequest(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusBadRequest, message))
}

func (ctx *Context) Conflict(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusConflict, message))
}

func (ctx *Context) Internal(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusInternalServerError, message))
}

func (ctx *Context) NotFound(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusNotFound, message))
}

func (ctx *Context) UnprocessableEntity(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusUnprocessableEntity, message))
}

func (ctx *Context) Forbidden(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusForbidden, message))
}

func (ctx *Context) Unauthorized(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusUnauthorized, message))
}
