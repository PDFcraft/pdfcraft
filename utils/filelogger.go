package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type FileStatStruct struct {
	CurrentTime time.Time
	FileName    string
	TimeDiff    int32
}

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func FileDeleteLogger() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	currentTime := time.Now()
	inFiles, _ := ioutil.ReadDir("./files/input")
	for _, inFile := range inFiles {
		timeDiff := int32(currentTime.Sub(inFile.ModTime()).Minutes())
		log.Println(inFile.Name(), timeDiff, "ms Ago")
	}
	outFiles, _ := ioutil.ReadDir("./files/input")
	for _, outFile := range outFiles {
		timeDiff := int32(currentTime.Sub(outFile.ModTime()).Minutes())
		log.Println(outFile.Name(), timeDiff, "ms Ago")
	}

}

func FileRecvLogger(recevedFiles map[int]string) {
	for i := 0; i < len(recevedFiles); i++ {
		file, _ := os.Stat("./files/input/" + recevedFiles[i])
		log.Println(file.Name(), file.ModTime())
	}
}
