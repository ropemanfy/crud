package errors

var (
	ErrInternal         = new("internal error", "internal")
	ErrNotFound         = new("not found", "notFound")
	ErrAccessDenied     = new("access denied", "accessDenied")
	ErrBadRequest       = new("bad request", "BadRequest")
	ErrUnauthorized     = new("you are not authorized", "unauthorized")
	ErrWrongQueryParams = new("missing required parameter", "missingParam")
)
