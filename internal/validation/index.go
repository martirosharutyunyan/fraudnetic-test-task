package validate

import (
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/shared/request"
)

// Validate is validator
var validation = validator.New()

func Body[T any](next func(*request.Context, *T)) func(ctx *request.Context) {
	return func(ctx *request.Context) {
		p := new(T)
		err := jsoniter.Unmarshal(ctx.Body(), p)
		if err != nil {
			ctx.UnprocessableEntity(err.Error())
			return
		}

		err = validation.StructCtx(ctx, p)
		if err != nil {
			ctx.UnprocessableEntity(err.Error())
			return
		}

		next(ctx, p)
	}
}

func Query[T any](next func(*request.Context, *T)) func(ctx *request.Context) {
	return func(ctx *request.Context) {
		p := new(T)
		err := ctx.QueryParser(p)
		if err != nil {
			ctx.UnprocessableEntity(err.Error())
			return
		}

		err = validation.StructCtx(ctx, p)
		if err != nil {
			ctx.UnprocessableEntity(err.Error())
			return
		}

		next(ctx, p)
	}
}
