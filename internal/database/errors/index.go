package dbErrors

import (
	"errors"
	httpErrors "github.com/go-boilerplate/common-module/exceptions/http-errors"
	"github.com/lib/pq"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/ent"
)

func Parse(err error) error {
	if err == nil {
		return nil
	}

	if ent.IsNotFound(err) {
		return httpErrors.NewNotFoundError(err.Error())
	}

	if ent.IsConstraintError(err) {
		var pqError *pq.Error
		isPqError := errors.As(errors.Unwrap(err), &pqError)
		if !isPqError {
			return httpErrors.NewInternalServerError("constraint error isn't pqError type")
		}

		switch pqError.Code {
		case foreignKeyErrorCode:
			errMessage, exists := foreignKeyConstraintErrors[pqError.Constraint]

			if !exists {
				return httpErrors.NewInternalServerError("constraint error isn't configured in foreignKeyConstraintErrors map")
			}

			return httpErrors.NewNotFoundError(errMessage)
		case uniqueKeyErrorCode:
			errMessage, exists := uniqueConstraintErrors[pqError.Constraint]

			if !exists {
				return httpErrors.NewInternalServerError("constraint error isn't configured in foreignKeyConstraintErrors map")
			}

			return httpErrors.NewConflictError(errMessage)
		}
	}

	return httpErrors.NewInternalServerError(err.Error())
}
