package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/javaandfly/gwebz/api/middleware"
	"github.com/javaandfly/gwebz/api/response"
)

func ParseJsonProcess(context *gin.Context, req any) {
	if err := context.ShouldBindWith(req, binding.JSON); err != nil {
		context.JSON(http.StatusBadRequest, response.ErrParam)
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// validate param
	err := middleware.ValidateParam(context, req)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.ErrParam)
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
}

func ParseFormProcess(context *gin.Context, req any) {
	if err := context.ShouldBindQuery(req); err != nil {
		context.JSON(http.StatusBadRequest, response.ErrParam)
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// validate param
	if err := middleware.ValidateParam(context, req); err != nil {
		context.JSON(http.StatusBadRequest, response.ErrParam)
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

}
