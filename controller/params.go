package controller

/*
 * @Descripttion: Store params for Gin http request
 */

// PublishActionParam
type PublishActionParam struct {
	Token string `json:"token,omitempty"`
	Data  []byte `json:"data,omitempty"`
	Title string `json:"title,omitempty"`
}

// PublishListParam
type PublishListParam struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

// FeedParam
type FeedParam struct {
	Token      string `json:"token" form:"token"`
	LatestTime int64  `json:"latest_time" form:"latest_time"`
}

// UserProfileParam
type UserProfileParam struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

// UserRegisterParam
type UserRegisterParam struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// UserLoginParam
type UserLoginParam struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
