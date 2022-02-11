package main

import (
	"github.com/PDFcraft/pdfcraft/pdfs"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/merge", pdfs.MergeHandler)

	router.GET("/download=:fileid", pdfs.Download)

	router.Run(":8080")
}
