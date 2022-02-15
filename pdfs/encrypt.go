package pdfs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func FileEncryptHandler(c *gin.Context) {
	password, uuidOrder, processedUuidName, fileNameDict := CommonHandler(c, "-lcoked")
	encryptPdfFile(uuidOrder[0], password, processedUuidName)
	c.JSON(http.StatusOK, gin.H{
		"FileName": fileNameDict,
	})
}

func encryptPdfFile(uuidFileName string, filePassword string, processedFileName string) {
	config := pdfcpu.NewAESConfiguration(filePassword, filePassword, 256)
	api.EncryptFile("./files/input/"+uuidFileName, "./files/output/"+processedFileName, config)
}
