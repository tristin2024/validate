// author: jiazujiang
// date: 2023/6/15
package validate

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New() // 实例化验证器
)

func Struct(s interface{}) error {
	return validate.Struct(s)
}
