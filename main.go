package main

import (
	"github.com/PDFcraft/pdfcraft/pdfs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/api/filetest", pdfs.FileTestHandler)

	router.POST("/api/merge", pdfs.MergeHandler)

	router.GET("/api/download=:fileid", pdfs.Download)

	router.POST("/api/decrypt", pdfs.FileDecryptHandler)
	router.POST("/api/encrypt", pdfs.FileEncryptHandler)

	router.Run(":8080")
}
