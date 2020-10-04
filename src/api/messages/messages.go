package messages

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const Stored = "保存しました"
const Modified = "更新しました"
const Deleted = "削除しました"
const NotFound = "存在しません"
const Forbidden = "権限がありません"
const EmailAlreadyExists = "このメールアドレスでは登録できません"
const RequiredError = "必須項目がありません"
const GeneralError = "エラーが発生しました"

const (
	InvalidEmailOrPassword = "メールアドレス、またはパスワードが一致しません"
)

// ExtractValidationErrorMsg はバリデーションエラーからエラーメッセージを抽出します
func ExtractValidationErrorMsg(err error) string {
	for _, err := range err.(validator.ValidationErrors) {
		var msg string
		field := err.Field()
		tag := err.Tag()
		switch tag {
		case "required":
			msg = fmt.Sprintf("%sは必須です", field)
		case "email":
			msg = "メールアドレスが不正です"
		case "gte":
			msg = fmt.Sprintf("%sは%s文字以上で入力してください", field, err.Param())
		case "lte":
			msg = fmt.Sprintf("%sは%s文字以下で入力してください", field, err.Param())
		case "min":
			msg = fmt.Sprintf("%sは%s以上で入力してください", field, err.Param())
		case "max":
			msg = fmt.Sprintf("%sは%s以下で入力してください", field, err.Param())
		}
		return msg
	}
	return ""
}
