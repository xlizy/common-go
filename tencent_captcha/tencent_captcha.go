package tencent_captcha

import (
	captcha "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha/v20190722"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/xlizy/common-go/config"
	"github.com/xlizy/common-go/enums/common_error"
	"github.com/xlizy/common-go/response"
	"github.com/xlizy/common-go/zlog"
)

type RootConfig struct {
	Config TencentCaptcha `yaml:"tencent-captcha"`
}

type TencentCaptcha struct {
	SecretId  string            `yaml:"secretId"`
	SecretKey string            `yaml:"secretKey"`
	Endpoint  string            `yaml:"endpoint"`
	Scenes    map[string]Scenes `yaml:"scenes"`
}

type Scenes struct {
	CaptchaAppId uint64 `yaml:"captchaAppId"`
	AppSecretKey string `yaml:"appSecretKey"`
	CaptchaType  uint64 `yaml:"captchaType"`
}

type CheckReq struct {
	Scenes string
	Ticket string
	UserIp string
}

var credential *common.Credential
var client *captcha.Client

var cfg TencentCaptcha

func InitTencentCaptcha(rc RootConfig) {
	cfg = rc.Config
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential = common.NewCredential(cfg.SecretId, cfg.SecretKey)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = cfg.Endpoint
	// 实例化要请求产品的client对象,clientProfile是可选的
	c, err := captcha.NewClient(credential, "", cpf)
	if err != nil {
		zlog.Error("创建腾讯验证码客户端异常:{}", err.Error())
	}
	client = c
}

func Check(req CheckReq) response.Response {
	if config.AppEnv.Env != "PROD" {
		return response.Success("非正式环境,默认成功", nil)
	}
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := captcha.NewDescribeCaptchaMiniResultRequest()
	request.CaptchaType = common.Uint64Ptr(cfg.Scenes[req.Scenes].CaptchaType)
	request.Ticket = common.StringPtr(req.Ticket)
	request.UserIp = common.StringPtr(req.Ticket)
	request.CaptchaAppId = common.Uint64Ptr(cfg.Scenes[req.Scenes].CaptchaAppId)
	request.AppSecretKey = common.StringPtr(cfg.Scenes[req.Scenes].AppSecretKey)

	// 返回的resp是一个DescribeCaptchaMiniResultResponse的实例，与请求对象对应
	resp, err := client.DescribeCaptchaMiniResult(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		zlog.Error("An API error has returned:{}", err.Error())
		return response.Error(common_error.CALL_CAPTCHA_ERROR, nil)
	}
	if err != nil {
		zlog.Error("调用验证码服务异常:{}", err.Error())
		return response.Error(common_error.CALL_CAPTCHA_ERROR, nil)
	}
	captchaCode := resp.Response.CaptchaCode
	if *captchaCode != int64(1) {
		// 1     ticket verification succeeded     票据验证成功
		// 7     CaptchaAppId does not match     票据与验证码应用APPID不匹配
		// 8     ticket expired     票据超时
		// 10    ticket format error     票据格式不正确
		// 15    ticket decryption failed     票据解密失败
		// 16    CaptchaAppId wrong format     检查验证码应用APPID错误
		// 21    (1)ticket error     票据验证错误 (2)diff 一般是由于用户网络较差，导致前端自动容灾，而生成了容灾票据，业务侧可根据需要进行跳过或二次处理
		// 25    invalid ticket     无效票据
		// 26    system internal error     系统内部错误
		// 31    UnauthorizedOperation.Unauthorized   无有效套餐包/账户已欠费
		// 100   param err     参数校验错误
		zlog.Error("行为验证失败:code:{},msg:{}", *captchaCode, *resp.Response.CaptchaMsg)
		return response.Error(common_error.SMART_VERIFICATION_ERROR, nil)
	} else {
		return response.Success("", nil)
	}
}
