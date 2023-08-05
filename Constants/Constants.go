package Constants

// privilege constants
const (
	USER_STATE_EXCEPTION  = "userStateException"
	PRIVILEGE_NOT_ALLOWED = "userPrivilegeNotAllowed"
)

// login constants
const (
	NOT_LOGIN                = "notLogin"
	LOGIN_SUCCESS            = "loginSuccess"
	LOGIN_FAILED             = "loginFailed"
	REQUEST_PARAMS_ERROR     = "RequestParamsError"
	USER_PROFILE_ALREAD_UESD = "userProfileAlreadyUSed"
)

// business constants
const (
	PULL_FAILED  = "pullFailure"
	PARAMS_ERROR = "paramsError"
	QUERY_FAILED = "queryFailed"
	MISMATCH     = "mismatch"
)

// db constants
const (
	DB_QUERY_FAILED  = "queryFailed"
	DB_SAVE_FAILED   = "saveFailed"
	DB_DELETE_FAILED = "deleteFailed"
	DB_MISMATCH      = "dataMismatch"
)

// general constants
const SUCCESS = "success"

const (
	User_NORMAL = "normalUser"
	Issuer      = "Auth_server"
)
