package rest_err

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string
	Message string
}

func NewRestError(message string, err string, code int, causes []Causes) RestErr {
	return	RestErr{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}
