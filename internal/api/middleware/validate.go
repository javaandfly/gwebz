package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// instantiation validate
func ValidateParamMiddleware(ctx *gin.Context) {

	validate := validator.New()
	ctx.Set("validate", validate)

	ctx.Next()

}

func ValidateParam(context *gin.Context, s interface{}) error {
	vali, isExist := context.Get("validate")
	if !isExist {
		return errors.New("validate context get error")
	}
	validate, ok := vali.(*validator.Validate)
	if !ok {
		return errors.New("validate context get error")
	}
	errs := validate.Struct(s)
	if errs != nil {
		return errs
	}
	return nil

}
