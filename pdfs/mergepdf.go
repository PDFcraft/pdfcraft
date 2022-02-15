package pdfs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergeHandler(c *gin.Context) {
	_, uuidOrder, processedUuidName, fileNameDict := CommonHandler(c, "-merged")
	mergePdfFile(uuidOrder, processedUuidName)
	c.JSON(http.StatusOK, gin.H{
		"FileName": fileNameDict,
	})
}

func mergePdfFile(recvFiles map[int]string, mergedFileName string) {
	inFiles := []string{}
	for i := 0; i < len(recvFiles); i++ {
		inFiles = append(inFiles, "./files/input/"+recvFiles[i])
	}
	api.MergeCreateFile(inFiles, "./files/output/"+mergedFileName, nil)
}
