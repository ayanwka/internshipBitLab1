package apperrors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomErrMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			lastErr := c.Errors.Last().Err
			var valErr ValidationError
			var nfErr NotFoundError
			var conErr ConflictError
			if errors.As(lastErr, &valErr) {
				c.Abort() // Прерываем цепочку Gin
				c.JSON(http.StatusBadRequest, gin.H{"error": valErr.Error()})
				return
			} else if errors.As(lastErr, &nfErr) {
				c.Abort() // Прерываем цепочку Gin
				c.JSON(http.StatusNotFound, gin.H{"error": nfErr.Error()})
				return
			} else if errors.As(lastErr, &conErr) {
				c.Abort() // Прерываем цепочку Gin
				c.JSON(http.StatusConflict, gin.H{"error": conErr.Error()})
				return
			} else {
				c.Abort() // Прерываем цепочку Gin
				c.JSON(500, gin.H{"error": "Internal Server Error"})
				return
			}
		}
	}
}
