package pdfs

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FileTestHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")
	c.Next()

	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	originName := make(map[string]string)
	extension := filepath.Ext(file.Filename)
	originFileName := filepath.Base(file.Filename)
	newFileName := uuid.New().String() + extension
	originName[originFileName] = newFileName
	if err := c.SaveUploadedFile(file, "./temp/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"FileName": originName,
	})

}
