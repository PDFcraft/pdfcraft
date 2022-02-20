package pdfs

import (
	"image"
	"net/http"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/PDFcraft/pdfcraft/utils"
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
	utils.FileProcessedLogger(processedUuidName, "IMG2PDF")
}
func imgToPdf(uuidImgOrder map[int]string, processedUuidName string) {
	for i := 0; i < len(uuidImgOrder); i++ {
		imgFile, _ := os.Open("./files/input/" + uuidImgOrder[i])
		imageConfig, _, _ := image.DecodeConfig(imgFile)
		if imageConfig.Width >= imageConfig.Height {
			imp, _ := api.Import("form:A4L, pos:c, sc:0.7 rel", pdfcpu.POINTS)
			api.ImportImagesFile([]string{"./files/input/" + uuidImgOrder[i]}, "./files/output/"+processedUuidName, imp, nil)
		} else {
			imp, _ := api.Import("form:A4P, pos:c, sc:0.7 rel", pdfcpu.POINTS)
			api.ImportImagesFile([]string{"./files/input/" + uuidImgOrder[i]}, "./files/output/"+processedUuidName, imp, nil)
		}
	}

	// api.RotateFile("./files/output/"+processedUuidName, "", 180, nil, nil)
}
