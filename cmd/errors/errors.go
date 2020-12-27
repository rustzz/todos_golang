package errors

type Errors struct {
	DATA_EMPTY_ERROR         map[string]interface{}
	PASSWORD_NOT_VALID_ERROR map[string]interface{}
	TOKEN_NOT_VALID_ERROR    map[string]interface{}
	USER_NOT_EXISTS_ERROR    map[string]interface{}
	USER_EXISTS_ERROR        map[string]interface{}
	NOTEBOOK_DATA_NOT_VALID  map[string]interface{}
}

func GetErrorsData() *Errors {
	return &Errors{
		DATA_EMPTY_ERROR: map[string]interface{}{
			"ok": false, "message": "Wrong request, recheck API docs"},

		PASSWORD_NOT_VALID_ERROR: map[string]interface{}{
			"ok": false, "message": "Password not valid"},

		USER_NOT_EXISTS_ERROR: map[string]interface{}{
			"ok": false, "message": "User not exists"},

		USER_EXISTS_ERROR: map[string]interface{}{
			"ok": false, "message": "User exists"},

		NOTEBOOK_DATA_NOT_VALID: map[string]interface{}{
			"ok": false, "message": "Wrong request, recheck API docs"},
	}
}
