package db

import (
	"fmt"

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
			_, err := t.CreateBucketIfNotExists([]byte("filename"))
			utils.HandleErr(err)
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

func SaveFileNamePair(uuidFileName string, originFileName string) {
	fmt.Println("Here : SaveFileNamePair")
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte("filename"))
		err := bucket.Put([]byte(uuidFileName), []byte(originFileName))
		return err
	})
	utils.HandleErr(err)
}

func GetFileNamePair(uuidFileName string) string {
	var originFileName string
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte("filename"))
		originFileName = string(bucket.Get([]byte(uuidFileName)))
		return nil
	})
	utils.HandleErr(err)

	return originFileName
}
