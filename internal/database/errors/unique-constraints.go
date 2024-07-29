package dbErrors

import "github.com/lib/pq"

var uniqueConstraintErrors = map[string]string{}

var uniqueKeyErrorCode pq.ErrorCode = "23505"
