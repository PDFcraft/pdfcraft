package utils

import (
	"fmt"

	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergePdfFile(recvFiles map[int]string) {
	inFiles := []string{}
	for _, filename := range recvFiles {
		inFiles = append(inFiles, "./temp/"+filename)
	}
	fmt.Print(inFiles)
	pdfcpu.MergeCreateFile(inFiles, "./output/merged.pdf", nil)
}
