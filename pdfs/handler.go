package pdfs

import (
	"net/http"
	"path/filepath"

	"github.com/PDFcraft/pdfcraft/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CommonHandler(c *gin.Context, feat string) (string, map[int]string, string, map[string]string) {
	password := c.PostForm("options")
	form, err := c.MultipartForm()
	files := form.File["files"] //[]*multipart.FileHeader
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
	}
	fileOrder := make(map[int]string)
	uuidOrder := make(map[int]string)
	fileNameDict := make(map[string]string)
	processedUuidName := uuid.New().String() + ".pdf"
	for i, file := range files {
		fileExtension := filepath.Ext(file.Filename)
		originFileName := filepath.Base(file.Filename)
		uuidFileName := uuid.New().String() + fileExtension
		fileOrder[i] = originFileName
		uuidOrder[i] = uuidFileName
		if err := c.SaveUploadedFile(file, "./files/input/"+uuidFileName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
		}
	}
	var processedFileName = fileOrder[0][0:len(fileOrder[0])-4] + feat + ".pdf"
	fileNameDict[processedFileName] = processedUuidName
	db.SaveFileNamePair(processedUuidName, processedFileName)

	return password, uuidOrder, processedUuidName, fileNameDict

}
