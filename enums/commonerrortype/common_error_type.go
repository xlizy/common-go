package common_error_type

type CommonErrorType int

const (
	SYSTEM_ERROR                  CommonErrorType = 1000
	SYSTEM_BUSY                   CommonErrorType = 1001
	FREQUENT_OPERATIONS           CommonErrorType = 1002
	TOO_MANY_MISTAKES             CommonErrorType = 1003
	VER_CODE_ERROR                CommonErrorType = 1004
	VER_CODE_NOT_EXIST            CommonErrorType = 1005
	VER_CODE_EXPIRED              CommonErrorType = 1006
	SMART_VERIFICATION_ERROR      CommonErrorType = 1007
	PARAMETER_MISSING             CommonErrorType = 1008
	VER_CODE_IS_USED              CommonErrorType = 1009
	REQUEST_ARGUMENT_NOT_VALID    CommonErrorType = 1010
	REQUEST_ARGUMENT_PARSER_ERROR CommonErrorType = 1011
	REQUEST_ARGUMENT_MISSING      CommonErrorType = 1012
	DUPLICATE_KEY                 CommonErrorType = -1013
	UPLOAD_FILE_SIZE_LIMIT        CommonErrorType = 1014
	UPLOAD_FILE_ERROR             CommonErrorType = 1015
	CHECK_DATA_SIGN_ERROR         CommonErrorType = 1016
	DECRYPT_DATA_ERROR            CommonErrorType = 1017
	DUBBO_SERVICE_UNAVAILABLE     CommonErrorType = 1018
	URL_NOT_FOUND                 CommonErrorType = 1019
	NOT_LOGGED_IN                 CommonErrorType = 1020
	SYS_ERR_ENUM_ERROR            CommonErrorType = -1021
	DEFAULT                       CommonErrorType = -1
)

func (e CommonErrorType) Code() int32 {
	return int32(e)
}
func (e CommonErrorType) Des() string {
	switch e {
	case SYSTEM_ERROR:
		return "通用错误"
	case SYSTEM_BUSY:
		return "系统繁忙,请稍候再试"
	case FREQUENT_OPERATIONS:
		return "操作频繁,请稍后再试"
	case TOO_MANY_MISTAKES:
		return "错误尝试过多,请稍后再试"
	case VER_CODE_ERROR:
		return "验证码错误"
	case VER_CODE_NOT_EXIST:
		return "验证码不存在"
	case VER_CODE_EXPIRED:
		return "验证码已过期"
	case SMART_VERIFICATION_ERROR:
		return "智能验证失败"
	case PARAMETER_MISSING:
		return "参数缺失"
	case VER_CODE_IS_USED:
		return "验证码已使用"
	case REQUEST_ARGUMENT_NOT_VALID:
		return "请求参数校验不通过"
	case REQUEST_ARGUMENT_PARSER_ERROR:
		return "请求参数解析失败"
	case REQUEST_ARGUMENT_MISSING:
		return "请求参数缺失"
	case DUPLICATE_KEY:
		return "主键或唯一键约束失败"
	case UPLOAD_FILE_SIZE_LIMIT:
		return "上传文件大小超过限制"
	case UPLOAD_FILE_ERROR:
		return "上传文件错误"
	case CHECK_DATA_SIGN_ERROR:
		return "数据验签失败"
	case DECRYPT_DATA_ERROR:
		return "数据解密失败"
	case DUBBO_SERVICE_UNAVAILABLE:
		return "系统繁忙,请稍候再试"
	case URL_NOT_FOUND:
		return "未匹配到请求URL"
	case NOT_LOGGED_IN:
		return "登录状态异常,请重新登录后操作"
	case SYS_ERR_ENUM_ERROR:
		return "response枚举转换错误"
	case DEFAULT:
		return "默认"
	}
	return "默认"
}
