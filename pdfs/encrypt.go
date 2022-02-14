package pdfs

import (
	"net/http"
	"path/filepath"

	"github.com/PDFcraft/pdfcraft/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func FileEncryptHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")
	c.Next()
	password := c.PostForm("options")
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	fileNameDict := make(map[string]string)
	fileExtension := filepath.Ext(file.Filename)
	originFileName := filepath.Base(file.Filename)
	uuidFileName := uuid.New().String() + fileExtension
	processedUuidName := uuid.New().String() + ".pdf"

	if err := c.SaveUploadedFile(file, "./files/input/"+uuidFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	encryptPdfFile(uuidFileName, password, processedUuidName)
	var processedFileName = originFileName[0:len(originFileName)-4] + "-locked" + ".pdf"
	fileNameDict[processedFileName] = processedUuidName
	db.SaveFileNamePair(processedUuidName, processedFileName)

	c.JSON(http.StatusOK, gin.H{
		"FileName": fileNameDict,
	})

}

func encryptPdfFile(uuidFileName string, filePassword string, processedFileName string) {
	config := pdfcpu.NewAESConfiguration(filePassword, filePassword, 256)
	api.EncryptFile("./files/input/"+uuidFileName, "./files/output/"+processedFileName, config)
}
