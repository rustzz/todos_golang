package errors

// Errors : ...
type Errors struct {
	DataEmptyError            map[string]interface{}
	PasswordNotValidError     map[string]interface{}
	TokenNotValidError        map[string]interface{}
	UserNotExistsError        map[string]interface{}
	UserExistsError           map[string]interface{}
	NotebookDataNotValidError map[string]interface{}
	RateLimitError            map[string]interface{}
	NoteNotFoundError         map[string]interface{}
}

// GetErrorsData : ...
func GetErrorsData() *Errors {
	return &Errors{
		DataEmptyError: map[string]interface{}{
			"ok": false, "message": "Wrong request, recheck API docs",
		},

		PasswordNotValidError: map[string]interface{}{
			"ok": false, "message": "Password not valid",
		},

		UserNotExistsError: map[string]interface{}{
			"ok": false, "message": "User not exists",
		},

		UserExistsError: map[string]interface{}{
			"ok": false, "message": "User exists",
		},

		NotebookDataNotValidError: map[string]interface{}{
			"ok": false, "message": "Wrong request, recheck API docs",
		},

		TokenNotValidError: map[string]interface{}{
			"ok": false, "message": "Token not valid",
		},
		RateLimitError: map[string]interface{}{
			"ok": false, "message": "You have reached maximum request limit",
		},
		NoteNotFoundError: map[string]interface{}{
			"ok": false, "message": "Note with that id not found",
		},
	}
}
