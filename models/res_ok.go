package models

type ResOK struct {
	Status int `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
