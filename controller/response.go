package controller

type response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func ok(code int, data interface{}) response {
	return response{
		Code:  code,
		Msg:   "操作成功",
		Data:  data,
		Error: "",
	}
}
