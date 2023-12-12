package request

import (
	"net/http"

	"ccian.cc/Satori/journey-src/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/javaandfly/gwebz/internal/api/response"
)

func ParseJsonProcess(context *gin.Context, req any) error {
	if err := context.ShouldBindWith(req, binding.JSON); err != nil {
		context.JSON(http.StatusBadRequest, response.ErrParam)
		return err
	}

	// validate param
	err := middleware.ValidateParam(context, req)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.ErrParam)
		return err
	}
	return nil
}

func ParseFormProcess(context *gin.Context, req any) error {
	if err := context.ShouldBindQuery(req); err != nil {
		context.JSON(http.StatusBadRequest, response.ErrParam)
		return err
	}

	// validate param
	if err := middleware.ValidateParam(context, req); err != nil {
		context.JSON(http.StatusBadRequest, response.ErrParam)
		return err
	}
	return nil

}
