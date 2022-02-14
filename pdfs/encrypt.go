package pdfs

import (
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func FileEncryptHandler(c *gin.Context) {
	password, uuidOrder, processedUuidName := CommonHandler(c, "-lcoked")
	encryptPdfFile(uuidOrder[0], password, processedUuidName)
}

func encryptPdfFile(uuidFileName string, filePassword string, processedFileName string) {
	config := pdfcpu.NewAESConfiguration(filePassword, filePassword, 256)
	api.EncryptFile("./files/input/"+uuidFileName, "./files/output/"+processedFileName, config)
}
