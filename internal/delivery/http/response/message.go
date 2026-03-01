package response

type MessageResponse struct {
	Message string `json:"message"`
}

func Message(message string) MessageResponse {
	return MessageResponse{
		Message: message,
	}
}
