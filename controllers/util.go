package controllers

type (
	empty struct{}

	response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func formatResponse(message string, data ...interface{}) *response {
	var datum interface{}
	datum = empty{}

	if message == "" {
		message = "Failed"
	}

	if len(data) > 0 {
		datum = data[0]
	}

	return &response{
		Message: message,
		Data:    datum,
	}
}
