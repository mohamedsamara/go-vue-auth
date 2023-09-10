package constants

type ContextKey string

const (
	API_URL              = "/api/v1"
	API_PROCESSING_ERROR = "Your request could not be processed. Please try again."
	AUTH_CONTEXT         = ContextKey("user")
	DB_CONTEXT           = ContextKey("db")
	ACCESS_TOKEN_EXPIRE  = 15
	REFRESH_TOKEN_EXPIRE = 60
)
