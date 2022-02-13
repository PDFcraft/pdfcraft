package pdfs

import (
	"net/http"
	"path/filepath"

	"github.com/PDFcraft/pdfcraft/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergeHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")
	c.Next()

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

	fileOrder := make(map[int]string)
	recvFiles := make(map[int]string)
	originName := make(map[string]string)
	mergedName := make(map[string]string)
	for i, file := range files {
		extension := filepath.Ext(file.Filename)
		originFileName := filepath.Base(file.Filename)
		newFileName := uuid.New().String() + extension
		fileOrder[i] = originFileName
		recvFiles[i] = newFileName
		originName[newFileName] = originFileName
		if err := c.SaveUploadedFile(file, "./files/input/"+newFileName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}

	}
	var mergedFileName = originName[recvFiles[0]][0:len(originName[recvFiles[0]])-4] + "-merged" + ".pdf"
	var newMergedName = uuid.New().String() + ".pdf"
	mergePdfFile(recvFiles, newMergedName)
	mergedName[mergedFileName] = newMergedName
	db.SaveFileNamePair(newMergedName, mergedFileName)
	c.JSON(http.StatusOK, gin.H{
		"FileName": mergedName,
	})
}

func mergePdfFile(recvFiles map[int]string, mergedFileName string) {
	inFiles := []string{}
	for i := 0; i < len(recvFiles); i++ {
		inFiles = append(inFiles, "./files/input/"+recvFiles[i])
	}
	pdfcpu.MergeCreateFile(inFiles, "./files/output/"+mergedFileName, nil)
}
