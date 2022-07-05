package products

func errResponse(errors ...error) ErrorResponse {
	var resp ErrorResponse
	resp.Errors = make([]ErrorMessage, len(errors))

	for i, err := range errors {
		resp.Errors[i] = ErrorMessage{err.Error()}
	}

	return resp
}
