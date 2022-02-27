package main

import (
	"github.com/PDFcraft/pdfcraft/pdfs"
	"github.com/PDFcraft/pdfcraft/utils"
	"github.com/jasonlvhit/gocron"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/merge", pdfs.MergeHandler)

	router.GET("/api/download=:fileid", pdfs.Download)

	router.POST("/api/unlock", pdfs.FileDecryptHandler)
	router.POST("/api/protect", pdfs.FileEncryptHandler)
	router.POST("/api/topdf", pdfs.ImgConvertHandler)
	router.POST("/api/split", pdfs.SplitHandler)

	go func() {
		s := gocron.NewScheduler()
		s.Every(1).Minutes().Do(utils.FileDeleteLogger)
		<-s.Start()
	}()

	router.Run(":8080")
}
