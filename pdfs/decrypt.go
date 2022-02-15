package pdfs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func FileDecryptHandler(c *gin.Context) {
	password, uuidOrder, processedUuidName, fileNameDict := CommonHandler(c, "-unlcoked")
	err := decryptPdfFile(uuidOrder[0], password, processedUuidName)
	if err != nil {
		errMsg := fmt.Sprintf("%s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"FileName": fileNameDict,
		})
	}
}
func decryptPdfFile(uuidFileName string, filePassword string, processedFileName string) error {
	config := pdfcpu.NewAESConfiguration(filePassword, filePassword, 256)
	err := api.DecryptFile("./files/input/"+uuidFileName, "./files/output/"+processedFileName, config)
	return err
}
