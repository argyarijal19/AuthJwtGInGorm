package middleware

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, statusCode int, statusMessage, errorMessage string) {
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"success": false,
		"status":  statusMessage,
		"data":    errorMessage,
	})
}

func SuccessResponse(c *gin.Context, statusCode int, statusMessage, data interface{}) {
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"success": true,
		"status":  statusMessage,
		"data":    data,
	})
}
