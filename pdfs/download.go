package pdfs

import (
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	fileid := c.Param("fileid")
	c.FileAttachment("./output/"+fileid+".pdf", "merged.pdf")

}
