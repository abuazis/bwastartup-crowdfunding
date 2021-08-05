package model

const BASE_URL = "http://localhost:8080/"

type WebResponse struct {
	Code   uint32      `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
