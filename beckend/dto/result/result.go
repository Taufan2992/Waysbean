package dto

type SuccessResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CartSuccessResult struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
