package mapError

type Messages struct {
	InvalidArgument string
	NotFound        string
	Unauthorized    string
	ServerError     string
}

func NewMessages() Messages {
	return Messages{
		InvalidArgument: "Invalid argument",
		NotFound:        "Not found",
		Unauthorized:    "Unauthorized",
		ServerError:     "Server error",
	}
}
