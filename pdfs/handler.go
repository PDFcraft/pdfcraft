package pdfs

import (
	"net/http"
	"path/filepath"

	"github.com/PDFcraft/pdfcraft/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CommonHandler(c *gin.Context, feat string) (string, map[int]string, map[int]string, string, map[string]string) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "POST,GET")
	c.Next()
	password := c.PostForm("options")
	form, err := c.MultipartForm()
	files := form.File["files"] //[]*multipart.FileHeader
	imgs := form.File["imgs"]
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
	}
	fileOrder := make(map[int]string)
	uuidOrder := make(map[int]string)
	fileNameDict := make(map[string]string)
	processedUuidName := uuid.New().String() + ".pdf"
	var processedFileName string
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
	imgOrder := make(map[int]string)
	imgUuidOrder := make(map[int]string)
	for j, img := range imgs {
		imgExtension := filepath.Ext(img.Filename)
		originImgName := filepath.Base(img.Filename)
		uuidImgName := uuid.New().String() + imgExtension
		imgOrder[j] = originImgName
		imgUuidOrder[j] = uuidImgName
		if err := c.SaveUploadedFile(img, "./files/input/"+uuidImgName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the img",
			})
		}
	}
	if len(fileOrder) > 1 {
		processedFileName = fileOrder[0][0:len(fileOrder[0])-4] + feat + ".pdf"
	} else {
		processedFileName = imgOrder[0][0:len(imgOrder[0])-4] + feat + ".pdf"
	}

	fileNameDict[processedFileName] = processedUuidName
	db.SaveFileNamePair(processedUuidName, processedFileName)

	return password, uuidOrder, imgUuidOrder, processedUuidName, fileNameDict

}
