package pdfs

import (
	"github.com/PDFcraft/pdfcraft/db"
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	fileid := c.Param("fileid")
	originFileName := db.GetFileNamePair(fileid + ".pdf")
	c.FileAttachment("./output/"+fileid+".pdf", originFileName)

}
