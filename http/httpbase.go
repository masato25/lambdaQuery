package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RespJSONError(c *gin.Context, msg string) {
	log.Println(msg)
	c.JSON(400, gin.H{
		"status": "error",
		"err":    msg,
	})
}

func RespJSON(c *gin.Context, msg gin.H) {
	c.JSON(200, msg)
}
