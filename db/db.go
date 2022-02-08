package db

import (
	"github.com/boltdb/bolt"
	"github.com/pdfcraft/pdfcraft/utils"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open("mergedfilename.db", 0600, nil)
		db = dbPointer
		utils.HandleErr(err)

	}
	return db
}
