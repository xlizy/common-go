package common_error

import (
	"github.com/xlizy/common-go/enums"
	"strconv"
)

type CommonError int

const (
	SYSTEM_ERROR                  CommonError = 1000
	SYSTEM_BUSY                   CommonError = 1001
	FREQUENT_OPERATIONS           CommonError = 1002
	TOO_MANY_MISTAKES             CommonError = 1003
	VER_CODE_ERROR                CommonError = 1004
	VER_CODE_NOT_EXIST            CommonError = 1005
	VER_CODE_EXPIRED              CommonError = 1006
	SMART_VERIFICATION_ERROR      CommonError = 1007
	PARAMETER_MISSING             CommonError = 1008
	VER_CODE_IS_USED              CommonError = 1009
	REQUEST_ARGUMENT_NOT_VALID    CommonError = 1010
	REQUEST_ARGUMENT_PARSER_ERROR CommonError = 1011
	REQUEST_ARGUMENT_MISSING      CommonError = 1012
	DUPLICATE_KEY                 CommonError = -1013
	UPLOAD_FILE_SIZE_LIMIT        CommonError = 1014
	UPLOAD_FILE_ERROR             CommonError = 1015
	CHECK_DATA_SIGN_ERROR         CommonError = 1016
	DECRYPT_DATA_ERROR            CommonError = 1017
	DUBBO_SERVICE_UNAVAILABLE     CommonError = 1018
	URL_NOT_FOUND                 CommonError = 1019
	NOT_LOGGED_IN                 CommonError = 1020
	SYS_ERR_ENUM_ERROR            CommonError = -1021
	ACCESS_DENIED                 CommonError = 1022
	CALL_CAPTCHA_ERROR            CommonError = 1023
	DEFAULT                       CommonError = -1
)

func (e CommonError) Code() int32 {
	return int32(e)
}
func (e CommonError) Des() string {
	return enums.BE(e).Des
}
func (e CommonError) MarshalJSON() ([]byte, error) {
	switch e {
	case SYSTEM_ERROR:
		return enums.JsonObj(int(e), "SYSTEM_ERROR", "通用错误")
	case SYSTEM_BUSY:
		return enums.JsonObj(int(e), "SYSTEM_BUSY", "系统繁忙,请稍候再试")
	case FREQUENT_OPERATIONS:
		return enums.JsonObj(int(e), "FREQUENT_OPERATIONS", "操作频繁,请稍后再试")
	case TOO_MANY_MISTAKES:
		return enums.JsonObj(int(e), "TOO_MANY_MISTAKES", "错误尝试过多,请稍后再试")
	case VER_CODE_ERROR:
		return enums.JsonObj(int(e), "VER_CODE_ERROR", "验证码错误")
	case VER_CODE_NOT_EXIST:
		return enums.JsonObj(int(e), "VER_CODE_NOT_EXIST", "验证码不存在")
	case VER_CODE_EXPIRED:
		return enums.JsonObj(int(e), "VER_CODE_EXPIRED", "验证码已过期")
	case SMART_VERIFICATION_ERROR:
		return enums.JsonObj(int(e), "SMART_VERIFICATION_ERROR", "行为验证失败")
	case PARAMETER_MISSING:
		return enums.JsonObj(int(e), "PARAMETER_MISSING", "参数缺失")
	case VER_CODE_IS_USED:
		return enums.JsonObj(int(e), "VER_CODE_IS_USED", "验证码已使用")
	case REQUEST_ARGUMENT_NOT_VALID:
		return enums.JsonObj(int(e), "REQUEST_ARGUMENT_NOT_VALID", "请求参数校验不通过")
	case REQUEST_ARGUMENT_PARSER_ERROR:
		return enums.JsonObj(int(e), "REQUEST_ARGUMENT_PARSER_ERROR", "请求参数解析失败")
	case REQUEST_ARGUMENT_MISSING:
		return enums.JsonObj(int(e), "REQUEST_ARGUMENT_MISSING", "请求参数缺失")
	case DUPLICATE_KEY:
		return enums.JsonObj(int(e), "DUPLICATE_KEY", "主键或唯一键约束失败")
	case UPLOAD_FILE_SIZE_LIMIT:
		return enums.JsonObj(int(e), "UPLOAD_FILE_SIZE_LIMIT", "上传文件大小超过限制")
	case UPLOAD_FILE_ERROR:
		return enums.JsonObj(int(e), "UPLOAD_FILE_ERROR", "上传文件错误")
	case CHECK_DATA_SIGN_ERROR:
		return enums.JsonObj(int(e), "CHECK_DATA_SIGN_ERROR", "数据验签失败")
	case DECRYPT_DATA_ERROR:
		return enums.JsonObj(int(e), "DECRYPT_DATA_ERROR", "数据解密失败")
	case DUBBO_SERVICE_UNAVAILABLE:
		return enums.JsonObj(int(e), "DUBBO_SERVICE_UNAVAILABLE", "系统繁忙,请稍候再试")
	case URL_NOT_FOUND:
		return enums.JsonObj(int(e), "URL_NOT_FOUND", "未匹配到请求URL")
	case NOT_LOGGED_IN:
		return enums.JsonObj(int(e), "NOT_LOGGED_IN", "登录状态异常,请重新登录后操作")
	case SYS_ERR_ENUM_ERROR:
		return enums.JsonObj(int(e), "SYS_ERR_ENUM_ERROR", "response枚举转换错误")
	case ACCESS_DENIED:
		return enums.JsonObj(int(e), "ACCESS_DENIED", "拒绝访问")
	case CALL_CAPTCHA_ERROR:
		return enums.JsonObj(int(e), "CALL_CAPTCHA_ERROR", "调用验证码服务异常")
	}
	return []byte(strconv.Itoa(int(e))), nil
}

func (e *CommonError) UnmarshalJSON(value []byte) error {
	*e = CommonError(enums.UnmarshalEnum(value))
	return nil
}
