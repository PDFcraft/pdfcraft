package utils

import (
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergePdfFile(inFiles []string) {
	pdfcpu.MergeCreateFile(inFiles, "out.pdf", nil)
}
