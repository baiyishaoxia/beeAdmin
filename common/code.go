package common


//状态码
const (
	Limit       = 15  //每页条数
	HttpSuccess = 200 //成功
	HttpError   = 201 //错误
	VueSuccess  = 1   //成功
	VueError    = 0   //错误
	VueRelogin  = 2   //重新登录
	PARAMS_ERROR             = 10001
	SYSTEM_ERROR             = 10002
	TOO_FREQUENTLY           = 10003
	QUERY_ERROR              = 10004
)
