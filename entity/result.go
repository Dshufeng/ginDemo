package entity

const (
	ERR_CODE = 400
	SUCCESS_CODE = 0
)
type Result struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (res *Result)SetCode(code int)  *Result{
	res.Code = code
	return res
}

func (res *Result)SetMsg(msg string) *Result {
	res.Msg = msg
	return res
}

func (res *Result)SetData(data interface{}) *Result {
	res.Data = data
	return res
}