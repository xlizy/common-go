package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	pv "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/kataras/iris/v12/context"
	"github.com/timandy/routine"
	"github.com/xlizy/common-go/enums/common_error"
	"github.com/xlizy/common-go/response"
	"reflect"
	"regexp"
)

const (
	ValidatorKey  = "ValidatorKey"
	TranslatorKey = "TranslatorKey"
	locale        = "chinese"
)

var validatorCtx = routine.NewInheritableThreadLocal()
var translatorCtx = routine.NewInheritableThreadLocal()

func TransInit(c *context.Context) {
	//设置支持语言
	chinese := zh.New()
	english := en.New()
	//设置国际化翻译器
	uni := ut.New(chinese, chinese, english)
	//设置验证器
	val := pv.New()
	//根据参数取翻译器实例
	trans, _ := uni.GetTranslator(locale)
	//翻译器注册到validator
	switch locale {
	case "chinese":
		zhTranslations.RegisterDefaultTranslations(val, trans)
		//使用fld.Tag.Get("comment")注册一个获取tag的自定义方法
		val.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})
		//自定义翻译器
		val.RegisterTranslation("email", trans, func(ut ut.Translator) error {
			return ut.Add("email", "{0}不符合邮箱规则", true)
		}, func(ut ut.Translator, fe pv.FieldError) string {
			t, _ := ut.T("email", fe.Field())
			return t
		})

		//自定义验证方法
		val.RegisterValidation("valid_username", func(fl pv.FieldLevel) bool {
			matched, _ := regexp.Match("^[a-z]{6,30}$", []byte(fl.Field().String()))
			return matched
		})
		//自定义翻译器
		val.RegisterTranslation("valid_username", trans, func(ut ut.Translator) error {
			return ut.Add("valid_username", "{0}输入格式不正确或长度不符", true)
		}, func(ut ut.Translator, fe pv.FieldError) string {
			t, _ := ut.T("valid_username", fe.Field())
			return t
		})

	case "english":
		enTranslations.RegisterDefaultTranslations(val, trans)
		val.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("en_comment")
		})
	}
	translatorCtx.Set(trans)
	validatorCtx.Set(val)
}

func ValidParams(params interface{}) response.Response {
	valid := validatorCtx.Get().(*pv.Validate)
	trans := translatorCtx.Get().(ut.Translator)
	err := valid.Struct(params)
	//如果数据效验不通过，则将所有err以切片形式输出
	if err != nil {
		errs := err.(pv.ValidationErrors)
		sliceErrs := make([]string, 0)
		for _, e := range errs {
			//使用validator.ValidationErrors类型里的Translate方法进行翻译
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return response.Error(common_error.REQUEST_ARGUMENT_NOT_VALID, sliceErrs)
	}
	return response.Success("", nil)
}
