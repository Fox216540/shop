package api

type Messages struct {
	MetadataNotFound       string
	UserIDMetadataNotFound string
	InvalidArgument        string
	NotFound               string
	Unauthorized           string
	Internal               string
	UserRegistered         string
	UserDeleted            string
	EmailUpdated           string
	PasswordUpdated        string
	PhoneUpdated           string
	ProfileUpdated         string
}

func NewMessages() Messages {
	return Messages{
		MetadataNotFound:       "metadata not found",
		UserIDMetadataNotFound: "user_id metadata not found",
		InvalidArgument:        "Invalid argument",
		NotFound:               "Not found",
		Unauthorized:           "Unauthorized",
		Internal:               "Internal server error",
		UserRegistered:         "User registered successfully",
		UserDeleted:            "User deleted successfully",
		EmailUpdated:           "Email updated successfully",
		PasswordUpdated:        "Password updated successfully",
		PhoneUpdated:           "Phone updated successfully",
		ProfileUpdated:         "Profile updated successfully",
	}
}

var messages = NewMessages()
