package db

import (
	"github.com/PDFcraft/pdfcraft/utils"
	"github.com/boltdb/bolt"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open("mergedfilename.db", 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte("fileName"))
			utils.HandleErr(err)
			return err
		})
		utils.HandleErr(err)

	}
	return db
}

func SaveFileNamePair(data []byte, uuidFileName string) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte("fileName"))
		err := bucket.Put([]byte(uuidFileName), data)
		return err
	})
	utils.HandleErr(err)
}
