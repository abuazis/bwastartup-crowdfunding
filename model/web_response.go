package model

type WebResponse struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data"`
}

type MetaResponse struct {
	Code    uint32 `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
