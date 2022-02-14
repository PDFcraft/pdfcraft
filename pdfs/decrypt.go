package pdfs

import (
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func FileDecryptHandler(c *gin.Context) {
	password, uuidOrder, processedUuidName := CommonHandler(c, "-unlcoked")
	decryptPdfFile(uuidOrder[0], password, processedUuidName)
}
func decryptPdfFile(uuidFileName string, filePassword string, processedFileName string) {
	config := pdfcpu.NewAESConfiguration(filePassword, filePassword, 256)
	api.DecryptFile("./files/input/"+uuidFileName, "./files/output/"+processedFileName, config)
}
