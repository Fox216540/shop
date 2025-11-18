package api

type Messages struct {
	LogInSuccess                  string
	LogOutSuccess                 string
	LogOutAllSuccess              string
	DeleteAllRefreshTokensSuccess string
}

func NewMessages() Messages {
	return Messages{
		LogInSuccess:                  "Log in success",
		LogOutSuccess:                 "Log out success",
		LogOutAllSuccess:              "Log out all success",
		DeleteAllRefreshTokensSuccess: "Delete all refresh tokens success",
	}
}
