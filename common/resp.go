package common

type Resp struct {
	ErrorCode interface{} `json:"error_code,omitempty"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data"`
	Paging    interface{} `json:"paging,omitempty"`
	Filter    interface{} `json:"filter,omitempty"`
}

func NewSuccessResp(data, paging, filter interface{}, msg string) *Resp {
	return &Resp{
		Data:      data,
		Paging:    paging,
		Filter:    filter,
		ErrorCode: 0,
		Message:   msg,
	}
}

func NewSimpleSuccessResp(data interface{}, msg string) *Resp {
	return &Resp{
		Data:      data,
		Paging:    nil,
		Filter:    nil,
		ErrorCode: nil,
		Message:   msg,
	}
}

func NewErrorResp(data, errorCode interface{}, msg string) *Resp {
	return &Resp{
		Data:      data,
		ErrorCode: errorCode,
	}
}
