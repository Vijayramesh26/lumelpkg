package common

const (
	SuccessCode = "S"
	ErrorCode   = "E"

	// Log levels
	INFO  = "INFO"
	ERROR = "ERROR"
	DEBUG = "DEBUG"

	// Database
	SELECT = "SELECT"
	INSERT = "INSERT"
	UPDATE = "UPDATE"

	DateLayout = "2006-01-02"
)

type CommonResp struct {
	DetailsArr any    `json:"respData"`
	Status     string `json:"status"`
	ErrMsg     string `json:"errMsg"`
}
