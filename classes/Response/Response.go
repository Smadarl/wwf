package Response

import "errors"

var (
	//Errors
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidUser = errors.New("invalid user")
	ErrUserNotFound = errors.New("user not found")
	ErrDataNotFound = errors.New("data not found")
	ErrModifiedToken = errors.New("token has been modified")
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("invalid token format/type")
	ErrCreateToken = errors.New("error creating token")
	ErrInternal = errors.New("internal error")
	ErrDataInput = errors.New("data input error")
	ErrRateLimit = errors.New("rate limit exceeded")
	ErrUserExists = errors.New("user exists")
	ErrMissingFields = errors.New("missing required fields")

	//Responses
	OK = &ApiResponse{Code: 200, Msg: "OK", Description: "Success"}
	NewResourceCreated = &ApiResponse{Code: 201, Msg: "Created", Description: "New resource created"}
	ResourceDeleted = &ApiResponse{Code: 204, Msg: "No Content", Description: "Resource deleted"}
)

type ApiResponse struct {
	Code        int `json:"code"`
	Msg         string `json:"message"`
	Description string `json:"description"`
}

func Error(err error) *ApiResponse {
	switch err {
	case ErrDataInput:
		return &ApiResponse{Code: 400, Msg: "Bad Request", Description: err.Error()}
	case ErrUserExists:
		return &ApiResponse{Code: 400, Msg: "Bad Request", Description: err.Error()}
	case ErrMissingFields:
		return &ApiResponse{Code: 400, Msg: "Bad Request", Description: err.Error()}
	case ErrUnauthorized:
		return &ApiResponse{Code: 401, Msg: "Unauthorized", Description: err.Error()}
	case ErrInvalidUser:
		return &ApiResponse{Code: 401, Msg: "Unauthorized", Description: err.Error()}
	case ErrModifiedToken:
		return &ApiResponse{Code: 401, Msg: "Unauthorized", Description: err.Error()}
	case ErrExpiredToken:
		return &ApiResponse{Code: 401, Msg: "Unauthorized", Description: err.Error()}
	case ErrInvalidToken:
		return &ApiResponse{Code: 401, Msg: "Unauthorized", Description: err.Error()}
	case ErrCreateToken:
		return &ApiResponse{Code: 401, Msg: "Unauthorized", Description: err.Error()}
	case ErrUserNotFound:
		return &ApiResponse{Code: 404, Msg: "Not Found", Description: err.Error()}
	case ErrDataNotFound:
		return &ApiResponse{Code: 404, Msg: "Not Found", Description: err.Error()}
	case ErrInternal: //Internal Server Error
		//TODO: handle internal errors
		return &ApiResponse{Code: 500, Msg: "Internal Server Error", Description: err.Error()}
	case ErrRateLimit: //server unavailable
		return &ApiResponse{Code: 503, Msg: "Server Unavailable", Description: err.Error()}
	}

	return nil
}