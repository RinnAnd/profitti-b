package dto

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
