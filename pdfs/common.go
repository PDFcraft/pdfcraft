package pdfs

import (
	"bytes"
	"encoding/gob"

	"github.com/PDFcraft/pdfcraft/db"
	"github.com/PDFcraft/pdfcraft/utils"
)

type FilePair struct {
	originFileName string `json:"originfilename"`
}

var f *FilePair

func (f *FilePair) toBytes() []byte {
	var blockBuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockBuffer)
	utils.HandleErr(encoder.Encode(f))
	return blockBuffer.Bytes()
}

func (f *FilePair) persist(uuidFileName string) {
	db.SaveFileNamePair(f.toBytes(), uuidFileName)
}
