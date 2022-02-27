package pdfs

import (
	"net/http"
	"strconv"

	"github.com/PDFcraft/pdfcraft/utils"
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func FileRotateHandler(c *gin.Context) {
	degree, uuidOrder, _, processedUuidName, fileNameDict := CommonHandler(c, "-rotated")
	intDegree, _ := strconv.Atoi(degree)
	rotatePdfFile(uuidOrder[0], intDegree, processedUuidName)
	c.JSON(http.StatusOK, gin.H{
		"FileName": fileNameDict,
	})
	utils.FileProcessedLogger(processedUuidName, "ROTATE")

}

func rotatePdfFile(uuidFileName string, degree int, processedFileName string) {
	api.RotateFile("./files/input/"+uuidFileName, "./files/output/"+processedFileName, degree, nil, nil)

}
