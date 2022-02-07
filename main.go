package main

import (
	"net/http"
	"path/filepath"
	"pdfcraft/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func saveFileHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	files := form.File["files"]

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	recvFiles := make(map[int]string)
	for i, file := range files {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		// originFileName := filepath.Base(file.Filename)

		recvFiles[i] = newFileName
		if err := c.SaveUploadedFile(file, "./temp/"+newFileName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}

	}
	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
		"files":   recvFiles,
	})

	utils.MergePdfFile(recvFiles)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/merge", saveFileHandler)
	router.GET("/merge", func(c *gin.Context) {
	})

	router.Run(":8080")
}
