package api

type Messages struct {
	LogInSuccess                  string
	LogOutSuccess                 string
	LogOutAllSuccess              string
	DeleteAllRefreshTokensSuccess string
	InvalidArgument               string
	NotFound                      string
	Unauthorized                  string
	ServerError                   string
}

func NewMessages() Messages {
	return Messages{
		LogInSuccess:                  "Log in success",
		LogOutSuccess:                 "Log out success",
		LogOutAllSuccess:              "Log out all success",
		DeleteAllRefreshTokensSuccess: "Delete all refresh tokens success",
		InvalidArgument:               "Invalid argument",
		NotFound:                      "Not found",
		Unauthorized:                  "Unauthorized",
		ServerError:                   "Server error",
	}
}
