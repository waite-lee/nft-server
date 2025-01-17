package mvc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary
// @Description 返回结果
type DataResult struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func Error(ctx *gin.Context, err error) {
	status := http.StatusInternalServerError
	ctx.JSON(status, &DataResult{
		Status:  status,
		Message: err.Error(),
		Data:    false,
		Success: false,
	})
}

func Fail(ctx *gin.Context, code int, message string) {
	status := http.StatusInternalServerError
	ctx.JSON(status, &DataResult{
		Status:  status,
		Message: message,
		Data:    false,
		Success: false,
	})
}

func Ok(ctx *gin.Context, data interface{}) {
	status := http.StatusOK
	ctx.JSON(status, &DataResult{
		Status:  status,
		Message: "",
		Data:    data,
		Success: true,
	})
}
