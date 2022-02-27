package pdfs

import (
	"errors"
	"os"

	"github.com/PDFcraft/pdfcraft/db"
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	fileid := c.Param("fileid")
	var processedFileName string
	var originFileName string
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")
	c.Next()
	if _, err := os.Stat("./files/output/" + fileid + ".pdf"); err == nil {
		processedFileName = "./files/output/" + fileid + ".pdf"
		originFileName = db.GetFileNamePair(fileid + ".pdf")
	} else if errors.Is(err, os.ErrNotExist) {
		processedFileName = "./files/output/" + fileid + ".zip"
		originFileName = db.GetFileNamePair(fileid) + ".zip"
	}
	c.FileAttachment(processedFileName, originFileName)

}
