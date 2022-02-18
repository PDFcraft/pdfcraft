package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	green   = "\033[97;42m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

func FileDeleteLogger() {
	iterate("./files")
}

func FileRecvLogger(recevedFiles map[int]string) {
	currentTime := time.Now()
	for i := 0; i < len(recevedFiles); i++ {
		file, _ := os.Stat("./files/input/" + recevedFiles[i])
		timeDiff := int32(currentTime.Sub(file.ModTime()).Minutes())
		fmt.Printf("[PDFCRAFT] |%s %-7s %s| %s | %5d %s |%s %-9s %s|\n", cyan, "INPUT", reset, file.Name(), timeDiff, "mS Ago", green, "RECIEVED", reset)
	}
}

func iterate(path string) {
	currentTime := time.Now()
	filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if !file.IsDir() && !(file.Name() == ".gitkeep") {
			timeDiff := int32(currentTime.Sub(file.ModTime()).Minutes())
			filepath := "OUTPUT"
			pathcolor := magenta
			status := "KEPT"
			statcolor := yellow
			if strings.Contains(path, "input") {
				pathcolor = cyan
				filepath = "INPUT"
			}
			if timeDiff > 60 {
				status = "DELETED"
				statcolor = red
			}
			fmt.Printf("[PDFCRAFT] |%s %-7s %s| %s | %5d %s |%s %-9s %s|\n", pathcolor, filepath, reset, file.Name(), timeDiff, "mS Ago", statcolor, status, reset)
			if timeDiff > 60 {
				os.Remove(path)
			}
		}
		return nil
	})
}
