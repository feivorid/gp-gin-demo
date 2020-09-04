package util

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(code int, message string, data interface{}) Response {
	var res Response
	res.Code = code
	res.Message = message
	res.Data = data
	return res
}
