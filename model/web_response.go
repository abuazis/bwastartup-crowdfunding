package model

type WebResponse struct {
	Code   uint32      `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
