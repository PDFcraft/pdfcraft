package pdfs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/PDFcraft/pdfcraft/utils"
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func SplitHandler(c *gin.Context) {
	var originFileName string
	_, uuidOrder, _, processedUuidName, fileNameDict := CommonHandler(c, "split")
	c.JSON(http.StatusOK, gin.H{
		"ProcessedName": fileNameDict,
	})
	for fileName, _ := range fileNameDict {
		originFileName = fileName
	}
	splitPdfFile(uuidOrder[0], processedUuidName, originFileName)
	utils.FileProcessedLogger(processedUuidName, "SPLITED")
}

func splitPdfFile(uuidFileName string, processedFileName string, fileName string) {
	os.Mkdir("./files/output/"+processedFileName, os.ModePerm)
	api.SplitFile("./files/input/"+uuidFileName, "./files/output/"+processedFileName, 1, nil)
	files, _ := ioutil.ReadDir("./files/output/" + processedFileName)
	fmt.Println(fileName)
	for i, file := range files {
		os.Rename("./files/output/"+processedFileName+"/"+file.Name(), "./files/output/"+processedFileName+"/"+fileName+"_"+strconv.Itoa(i+1)+".pdf")
	}
	utils.ZipSource("./files/output/"+processedFileName, "./files/output/"+processedFileName+".zip")
}
