package api

type Messages struct {
	LogInSuccess                  string
	LogOutSuccess                 string
	LogOutAllSuccess              string
	DeleteAllRefreshTokensSuccess string
	InvalidArgument               string
	InvalidJSON                   string
	NotFound                      string
	Unauthorized                  string
	MissingRefreshCookie          string
	MissingUserID                 string
	RefreshCookieName             string
	RefreshSuccess                string
	ServiceUnavailable            string
	ServerError                   string
}

func NewMessages() Messages {
	return Messages{
		LogInSuccess:                  "Log in success",
		LogOutSuccess:                 "Log out success",
		LogOutAllSuccess:              "Log out all success",
		DeleteAllRefreshTokensSuccess: "Delete all refresh tokens success",
		InvalidArgument:               "Invalid argument",
		InvalidJSON:                   "Invalid JSON payload",
		NotFound:                      "Not found",
		Unauthorized:                  "Unauthorized",
		MissingRefreshCookie:          "Missing refresh cookie",
		MissingUserID:                 "User id not found",
		RefreshCookieName:             "__Host-refresh",
		RefreshSuccess:                "Access and refresh tokens updated successfully",
		ServiceUnavailable:            "Service unavailable",
		ServerError:                   "Server error",
	}
}

var messages = NewMessages()
