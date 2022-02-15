package main

import (
	"time"

	"github.com/PDFcraft/pdfcraft/pdfs"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/api/merge", pdfs.MergeHandler)

	router.GET("/api/download=:fileid", pdfs.Download)

	router.POST("/api/unlock", pdfs.FileDecryptHandler)
	router.POST("/api/protect", pdfs.FileEncryptHandler)

	router.Run(":8080")
}
