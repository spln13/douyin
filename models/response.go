package models

type CommonResponseBody struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_msg"`
}
