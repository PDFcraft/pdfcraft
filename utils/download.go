package utils

import (
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	fileid := c.Param("fileid")
	c.File("./output/" + fileid + ".pdf")

}
