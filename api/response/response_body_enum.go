package response

var (
	// OK
	SUCCESS = response(200, "SUCCESS")
	Err     = response(500, "FAILED")

	// service level error code

	ErrParam   = response(10001, "invalid param")
	ErrDB      = response(10002, "db init error")
	ErrSQL     = response(10003, "db SQL error")
	ErrNotFind = response(10005, "not found")
	// ......
)
