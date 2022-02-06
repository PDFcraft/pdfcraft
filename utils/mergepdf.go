package utils

import (
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergePdfFile() {
	inFiles := []string{"in1.pdf", "in2.pdf"}
	pdfcpu.MergeCreateFile(inFiles, "out.pdf", nil)
}
