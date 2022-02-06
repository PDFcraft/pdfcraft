package main

import (
	"net/http"
	"path/filepath"
	"pdfcraft/utils"

	"github.com/gin-gonic/gin"
)

func handleFunc(c *gin.Context) {
	// file path in server
	c.File("out.pdf")
}

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	inFiles := []string{}

	router.POST("/merge", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}
		files := form.File["files"]
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			inFiles = append(inFiles, filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
				return
			}
		}
		utils.MergePdfFile(inFiles)
		c.HTML(http.StatusOK, "merge.html", gin.H{})

	})

	router.GET("/download", handleFunc)

	router.Run(":8080")
}
