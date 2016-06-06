package store

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var DB bolt.DB
var world = []byte("world")

func init() {
}

func AddToDb(bucket []byte, key []byte, value []byte) error {
	DB, err := bolt.Open("my.db", 0644, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()
	err = DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func RetrieveFromDb(bucket []byte, key []byte) error {
	db, err := bolt.Open("my.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", key)
		}

		val := bucket.Get(key)
		fmt.Println(string(val))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}