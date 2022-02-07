package main

import (
	"pdfcraft/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/merge", utils.MergeHandler)

	router.Run(":8080")
}
