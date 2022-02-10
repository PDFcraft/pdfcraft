package pdfs

import (
	"github.com/PDFcraft/pdfcraft/db"
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")
	c.Next()

	fileid := c.Param("fileid")
	originFileName := db.GetFileNamePair(fileid + ".pdf")
	c.FileAttachment("./output/"+fileid+".pdf", originFileName)

}
