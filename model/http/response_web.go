package model

type ResponseWebSuccess struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWebFailed struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type NotFoundError struct {
	Message string `json:"message"`
}

type UnauthorizedError struct {
	Message string `json:"message"`
}

type HttpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type BadRequestError struct {
	Message string `json:"message"`
}

type ConflictError struct {
	Message string `json:"message"`
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

func (e *ConflictError) Error() string {
	return e.Message
}

func (e *NotFoundError) Error() string {
	return e.Message
}
func (e *BadRequestError) Error() string {
	return e.Message
}
func (e *HttpError) Error() string {
	return e.Message
}
