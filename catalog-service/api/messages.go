package api

type Messages struct {
	MetadataNotFound       string
	UserIDMetadataNotFound string
	InvalidArgument        string
	NotFound               string
	Unauthorized           string
	Internal               string
}

func NewMessages() Messages {
	return Messages{
		MetadataNotFound:       "metadata not found",
		UserIDMetadataNotFound: "user_id metadata not found",
		InvalidArgument:        "Invalid argument",
		NotFound:               "Not found",
		Unauthorized:           "Unauthorized",
		Internal:               "Internal server error",
	}
}

var messages = NewMessages()
