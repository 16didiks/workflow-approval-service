package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, APIResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}

func Error(c *gin.Context, code int, err interface{}) {
	c.JSON(code, APIResponse{
		Success: false,
		Data:    nil,
		Error:   err,
	})
}
