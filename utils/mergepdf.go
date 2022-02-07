package utils

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergeHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	files := form.File["files"]

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	recvFiles := make(map[int]string)
	originName := make(map[string]string)
	for i, file := range files {
		extension := filepath.Ext(file.Filename)
		originFileName := filepath.Base(file.Filename)
		newFileName := uuid.New().String() + extension
		// Linking uuid with sent order with dict
		recvFiles[i] = newFileName
		originName[newFileName] = originFileName
		if err := c.SaveUploadedFile(file, "./temp/"+newFileName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}

	}
	// var name = filepath.Ext(originName[recvFiles[0]])
	var mergedFileName = originName[recvFiles[0]][0:len(originName[recvFiles[0]])-4] + "-merged" + ".pdf"
	// var name = TrimRight(originName[recvFiles[0]], ".pdf")
	mergePdfFile(recvFiles, mergedFileName)
	c.File("./output/" + mergedFileName)
}

func mergePdfFile(recvFiles map[int]string, mergedFileName string) {
	inFiles := []string{}
	for _, filename := range recvFiles {
		inFiles = append(inFiles, "./temp/"+filename)
	}
	pdfcpu.MergeCreateFile(inFiles, "./output/"+mergedFileName, nil)
}
