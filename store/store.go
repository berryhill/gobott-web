package store

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var DB bolt.DB
var world = []byte("world")

func init() {
	DB, err := bolt.Open("/home/Coding/gopath/src/github.com/gobott-web", 0644, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()
}

func AddToDb(bucket []byte, key []byte, value []byte) error {
	err := DB.Update(func(tx *bolt.Tx) error {
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

func RetrieveFromDb(input []byte, key []byte) error {
	err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", world)
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