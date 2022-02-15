package main

import (
	"github.com/PDFcraft/pdfcraft/pdfs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/merge", pdfs.MergeHandler)

	router.GET("/api/download=:fileid", pdfs.Download)

	router.POST("/api/unlock", pdfs.FileDecryptHandler)
	router.POST("/api/protect", pdfs.FileEncryptHandler)

	router.Run(":8080")

}
