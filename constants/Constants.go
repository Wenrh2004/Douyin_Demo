package constants

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

// status code
const (
	STATUS_SUCCESS      = 0
	STATUS_FAILED       = 1
	STATUS_UNABLE_QUERY = 300
	STATUS_INTERNAL_ERR = 301
	PARAMS_ERROR_CODE   = 302
)

// status msg
var (
	// common
	INTERNAL_SERVER_ERROR = "internalServerError, please try again later"
	// publish action
	SUCCESS_PUBLISH           = "publishSuccess"
	INVALID_CONTENT_TYPE      = "Invalid Content Type, only support mp4"
	UPLOAD_FAILED             = "uploadFailed, please try again later"
	GET_THUMBNAIL_LINK_FAILED = "create cover failed, please try again later"
)
