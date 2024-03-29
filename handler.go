package ginjson

import "github.com/gin-gonic/gin"
import "net/http"

type Handler func(c *Context) (any, error)

func (h Handler) ToGin() func(c *gin.Context) {
	return func(c *gin.Context) {
		resp, err := h(&Context{c})
		if err == nil {
			if resp != nil {
				c.JSON(http.StatusOK, resp)
				return
			}
		} else {
			switch e := err.(type) {
			case Error:
				c.JSON(e.Code, e.Response)
				return
			default: // fallback for generic error
				c.JSON(http.StatusInternalServerError, e.Error())
				return
			}
		}
	}
}
