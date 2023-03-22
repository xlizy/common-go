package response

type Response struct {
	Success      bool   `protobuf:"bool,1,opt,name=success,proto3" json:"success"`
	Code         int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code"`
	Msg          string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	ResponseTime string `protobuf:"bytes,4,opt,name=responseTime,proto3" json:"responseTime,omitempty"`
	TraceId      string `protobuf:"bytes,5,opt,name=traceId,proto3" json:"traceId,omitempty"`
}

type Responses interface {
	SetCode(int32)
	SetTraceID(string)
	SetMsg(string)
	SetData(interface{})
	SetSuccess(bool)
	SetResponseTime(string)
	Clone() Responses
}

type response struct {
	Response
	Data interface{} `protobuf:"bytes,6,opt,name=traceId,proto3" json:"data,omitempty"`
}

type Page struct {
	Count     int `json:"count"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type page struct {
	Page
	List interface{} `json:"list"`
}

func (e *response) SetData(data interface{}) {
	e.Data = data
}

func (e response) Clone() Responses {
	return &e
}

func (e *response) SetTraceID(id string) {
	e.TraceId = id
}

func (e *response) SetMsg(s string) {
	e.Msg = s
}

func (e *response) SetCode(code int32) {
	e.Code = code
}

func (e *response) SetResponseTime(responseTime string) {
	e.ResponseTime = responseTime
}

func (e *response) SetSuccess(success bool) {
	e.Success = success
	if !success {
		e.Success = false
	}
}
