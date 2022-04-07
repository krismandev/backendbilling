package core

// Internal Error message definition
const (
	//Success
	ErrSuccess  int    = 0
	DescSuccess string = "Success"

	ErrFailed  int    = 1
	DescFailed string = "Failed"

	ErrServer  int    = 9500
	DescServer string = "Server Error"

	ErrNotAuthorized  int    = 8123
	DescNotAuthorized string = "Not Authorized"

	ErrUnknown  int    = 3002
	DescUnknown string = "[3002]Unknown Error"

	ErrIncompleteRequest  int    = 3003
	DescIncompleteRequest string = "Incomplete Request"

	ErrNoData  int    = 6990
	DescNoData string = "No Data"

	ErrDataExists  int    = 6011
	DescDataExists string = "Data is already exists"

	ErrDeleteData  int    = 6012
	DescDeleteData string = "Cannot delete the data"

	//Query Error
	ErrFailedQuery  int    = 7010
	DescFailedQuery string = "Failed when process the database"

	//Redis Error
	ErrFailedRedis  int    = 7020
	DescFailedRedis string = "Redis Error"

	ErrTokenExpired  int    = 8131
	DescTokenExpired string = "Token Expired"

	ErrNoAccessRight  int    = 8150
	DescNoAccessRight string = "No Access Right"

	//Bad Request
	ErrOthers  int    = 9999
	DescOthers string = "[9999]Unknown Error"

	ErrInvalidFormat  int    = 9910
	DescInvalidFormat string = "Invalid request format"

	ErrRequestThrottled  int    = 2082
	DescRequestThrottled string = "Request Throttled"
)
