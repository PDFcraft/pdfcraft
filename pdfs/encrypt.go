package pdfs

import (
	"fmt"
	"net/http"

	"github.com/PDFcraft/pdfcraft/utils"
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func FileEncryptHandler(c *gin.Context) {
	password, uuidOrder, _, processedUuidName, fileNameDict := CommonHandler(c, "-lcoked")
	fmt.Println(uuidOrder)
	encryptPdfFile(uuidOrder[0], password, processedUuidName)
	c.JSON(http.StatusOK, gin.H{
		"FileName": fileNameDict,
	})
	utils.FileProcessedLogger(processedUuidName, "ENCRYPTED")

}

func encryptPdfFile(uuidFileName string, filePassword string, processedFileName string) {
	config := pdfcpu.NewAESConfiguration(filePassword, filePassword, 256)
	api.EncryptFile("./files/input/"+uuidFileName, "./files/output/"+processedFileName, config)
}
