package api

type Messages struct {
	InvalidArgument string
	NotFound        string
	Unauthorized    string
	Conflict        string
	Internal        string
}

func NewMessages() Messages {
	return Messages{
		InvalidArgument: "Invalid argument",
		NotFound:        "Not found",
		Unauthorized:    "Unauthorized",
		Conflict:        "Conflict",
		Internal:        "Internal server error",
	}
}

var messages = NewMessages()
