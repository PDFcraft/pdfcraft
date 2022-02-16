package pdfs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergeHandler(c *gin.Context) {
	_, uuidOrder, _, processedUuidName, fileNameDict := CommonHandler(c, "-merged")
	mergePdfFile(uuidOrder, processedUuidName)
	c.JSON(http.StatusOK, gin.H{
		"FileName": fileNameDict,
	})
}

func mergePdfFile(uuidOrder map[int]string, processedFileName string) {
	inFiles := []string{}
	for i := 0; i < len(uuidOrder); i++ {
		inFiles = append(inFiles, "./files/input/"+uuidOrder[i])
	}
	api.MergeCreateFile(inFiles, "./files/output/"+processedFileName, nil)
}
