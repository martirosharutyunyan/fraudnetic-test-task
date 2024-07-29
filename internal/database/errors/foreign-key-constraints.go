package dbErrors

import "github.com/lib/pq"

var foreignKeyConstraintErrors = map[string]string{}

var foreignKeyErrorCode pq.ErrorCode = "23503"
