package api

type Messages struct {
	MetadataNotFound          string
	UserIDMetadataNotFound    string
	InvalidArgument           string
	NotFound                  string
	Unauthorized              string
	ServerError               string
	BasketDeletedSuccessfully string
	ItemRemovedSuccessfully   string
}

func NewMessages() Messages {
	return Messages{
		MetadataNotFound:          "metadata not found",
		UserIDMetadataNotFound:    "user_id metadata not found",
		InvalidArgument:           "Invalid argument",
		NotFound:                  "Not found",
		Unauthorized:              "Unauthorized",
		ServerError:               "Server error",
		BasketDeletedSuccessfully: "Basket deleted successfully",
		ItemRemovedSuccessfully:   "Item removed successfully",
	}
}

var messages = NewMessages()
