package pdfs

import (
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergeHandler(c *gin.Context) {
	_, uuidOrder, processedUuidName := CommonHandler(c, "-merged")
	mergePdfFile(uuidOrder, processedUuidName)
}

func mergePdfFile(recvFiles map[int]string, mergedFileName string) {
	inFiles := []string{}
	for i := 0; i < len(recvFiles); i++ {
		inFiles = append(inFiles, "./files/input/"+recvFiles[i])
	}
	api.MergeCreateFile(inFiles, "./files/output/"+mergedFileName, nil)
}
