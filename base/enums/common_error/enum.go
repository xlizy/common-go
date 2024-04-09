package common_error

import (
	"github.com/xlizy/common-go/base/enums"
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
	HTTP_CALL_ERROR               CommonError = 1024
	HTTP_CALL_TIMEOUT             CommonError = 1025
	RPC_CALL_ERROR                CommonError = 1026
	MYSQL_ERR_1008                CommonError = 1030
	MYSQL_ERR_1012                CommonError = 1031
	MYSQL_ERR_1020                CommonError = 1032
	MYSQL_ERR_1021                CommonError = 1033
	MYSQL_ERR_1022                CommonError = 1034
	MYSQL_ERR_1037                CommonError = 1035
	MYSQL_ERR_1044                CommonError = 1036
	MYSQL_ERR_1045                CommonError = 1037
	MYSQL_ERR_1048                CommonError = 1038
	MYSQL_ERR_1049                CommonError = 1039
	MYSQL_ERR_1054                CommonError = 1040
	MYSQL_ERR_1062                CommonError = 1041
	MYSQL_ERR_1065                CommonError = 1042
	MYSQL_ERR_1114                CommonError = 1043
	MYSQL_ERR_1130                CommonError = 1044
	MYSQL_ERR_1133                CommonError = 1045
	MYSQL_ERR_1141                CommonError = 1046
	MYSQL_ERR_1142                CommonError = 1047
	MYSQL_ERR_1143                CommonError = 1048
	MYSQL_ERR_1149                CommonError = 1049
	MYSQL_ERR_1169                CommonError = 1051
	MYSQL_ERR_1216                CommonError = 1052
	INSUFFICIENT_SCOPE            CommonError = 1060
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
	case HTTP_CALL_ERROR:
		return enums.JsonObj(int(e), "HTTP_CALL_ERROR", "http请求失败")
	case HTTP_CALL_TIMEOUT:
		return enums.JsonObj(int(e), "HTTP_CALL_TIMEOUT", "http请求超时")
	case RPC_CALL_ERROR:
		return enums.JsonObj(int(e), "RPC_CALL_ERROR", "rpc请求失败")
	case MYSQL_ERR_1008:
		return enums.JsonObj(int(e), "MYSQL_ERR_1008", "数据库不存在")
	case MYSQL_ERR_1012:
		return enums.JsonObj(int(e), "MYSQL_ERR_1012", "不能读取系统表中的记录")
	case MYSQL_ERR_1020:
		return enums.JsonObj(int(e), "MYSQL_ERR_1020", "记录已被其他用户修改")
	case MYSQL_ERR_1021:
		return enums.JsonObj(int(e), "MYSQL_ERR_1021", "硬盘剩余空间不足，请加大硬盘可用空间")
	case MYSQL_ERR_1022:
		return enums.JsonObj(int(e), "MYSQL_ERR_1022", "关键字重复，更改记录失败")
	case MYSQL_ERR_1037:
		return enums.JsonObj(int(e), "MYSQL_ERR_1037", "系统内存不足，请重启数据库或重启服务器")
	case MYSQL_ERR_1044:
		return enums.JsonObj(int(e), "MYSQL_ERR_1044", "当前用户没有访问数据库的权限")
	case MYSQL_ERR_1045:
		return enums.JsonObj(int(e), "MYSQL_ERR_1045", "不能连接数据库，用户名或密码错误")
	case MYSQL_ERR_1048:
		return enums.JsonObj(int(e), "MYSQL_ERR_1048", "字段不能为空")
	case MYSQL_ERR_1049:
		return enums.JsonObj(int(e), "MYSQL_ERR_1049", "数据库不存在")
	case MYSQL_ERR_1054:
		return enums.JsonObj(int(e), "MYSQL_ERR_1054", "字段不存在")
	case MYSQL_ERR_1062:
		return enums.JsonObj(int(e), "MYSQL_ERR_1062", "违反唯一性约束")
	case MYSQL_ERR_1065:
		return enums.JsonObj(int(e), "MYSQL_ERR_1065", "无效的SQL语句，SQL语句为空")
	case MYSQL_ERR_1114:
		return enums.JsonObj(int(e), "MYSQL_ERR_1114", "数据表已满，不能容纳任何记录")
	case MYSQL_ERR_1130:
		return enums.JsonObj(int(e), "MYSQL_ERR_1130", "连接数据库失败，没有连接数据库的权限")
	case MYSQL_ERR_1133:
		return enums.JsonObj(int(e), "MYSQL_ERR_1133", "数据库用户不存在")
	case MYSQL_ERR_1141:
		return enums.JsonObj(int(e), "MYSQL_ERR_1141", "当前用户无权访问数据库")
	case MYSQL_ERR_1142:
		return enums.JsonObj(int(e), "MYSQL_ERR_1142", "当前用户无权访问数据表")
	case MYSQL_ERR_1143:
		return enums.JsonObj(int(e), "MYSQL_ERR_1143", "当前用户无权访问数据表中的字段")
	case MYSQL_ERR_1149:
		return enums.JsonObj(int(e), "MYSQL_ERR_1149", "SQL语句语法错误")
	case MYSQL_ERR_1169:
		return enums.JsonObj(int(e), "MYSQL_ERR_1169", "字段值重复，更新记录失败")
	case MYSQL_ERR_1216:
		return enums.JsonObj(int(e), "MYSQL_ERR_1216", "外键约束检查失败，更新子表记录失败")
	case INSUFFICIENT_SCOPE:
		return enums.JsonObj(int(e), "INSUFFICIENT_SCOPE", "授权不足")

	}
	return []byte(strconv.Itoa(int(e))), nil
}

func (e *CommonError) UnmarshalJSON(value []byte) error {
	*e = CommonError(enums.UnmarshalEnum(value))
	return nil
}
