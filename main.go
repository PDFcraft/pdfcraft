package main

import (
	"github.com/pdfcraft/pdfcraft/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/merge", utils.MergeHandler)

	router.GET("/download=:fileid", utils.Download)

	router.Run(":8080")
}
