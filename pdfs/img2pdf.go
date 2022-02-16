package pdfs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func ImgConvertHandler(c *gin.Context) {
	_, _, uuidImgOrder, processedUuidName, fileNameDict := CommonHandler(c, "-converted")
	imgToPdf(uuidImgOrder, processedUuidName)
	c.JSON(http.StatusOK, gin.H{
		"FileName": fileNameDict,
	})
}
func imgToPdf(uuidImgOrder map[int]string, processedUuidName string) {
	inFiles := []string{}
	for i := 0; i < len(uuidImgOrder); i++ {
		inFiles = append(inFiles, "./files/input/"+uuidImgOrder[i])
	}
	imp, _ := api.Import("form:A4, pos:c, s:1.0", pdfcpu.POINTS)
	api.ImportImagesFile(inFiles, "./files/output/"+processedUuidName, imp, nil)
	api.RotateFile("./files/output/"+processedUuidName, "", 180, nil, nil)
}
