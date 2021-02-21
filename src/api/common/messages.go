package common

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Message string

var (
	Stored                 Message = "保存しました"
	Modified               Message = "更新しました"
	Deleted                Message = "削除しました"
	NotFound               Message = "存在しません"
	Forbidden              Message = "権限がありません"
	EmailAlreadyExists     Message = "このメールアドレスでは登録できません"
	InvalidEmailOrPassword Message = "メールアドレス、またはパスワードが一致しません"
	InvalidRequest         Message = "リクエストが不正です"
	GeneralError           Message = "エラーが発生しました"
)

func (m Message) String() string {
	return string(m)
}

// ExtractValidationErrorMsg はバリデーションエラーからエラーメッセージを抽出します
func ExtractValidationErrorMsg(err error) string {
	for _, e := range err.(validator.ValidationErrors) {
		var msg string
		field := e.Field()
		tag := e.Tag()
		switch tag {
		case "required":
			msg = fmt.Sprintf("%sは必須です", field)
		case "email":
			msg = "メールアドレスが不正です"
		case "gte":
			msg = fmt.Sprintf("%sは%s文字以上で入力してください", field, e.Param())
		case "lte":
			msg = fmt.Sprintf("%sは%s文字以下で入力してください", field, e.Param())
		case "min":
			msg = fmt.Sprintf("%sは%s以上で入力してください", field, e.Param())
		case "max":
			msg = fmt.Sprintf("%sは%s以下で入力してください", field, e.Param())
		default:
			msg = err.Error()
		}
		return msg
	}
	return ""
}
