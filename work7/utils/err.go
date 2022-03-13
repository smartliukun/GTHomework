package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleErr(c *gin.Context, err *error) {
	c.JSON(http.StatusOK, gin.H{"code": 999,
		"msg": "handle request fail:err = " + (*err).Error(),
	})
}
