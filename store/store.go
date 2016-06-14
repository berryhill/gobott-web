package store

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"encoding/binary"
)

func openDb() (*bolt.DB, error) {
	db, err := bolt.Open("my.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))

	return b
}

func AddToDb(bucket []byte, value []byte) error {
	db, err := openDb()
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucket)

		if err != nil {
			return err
		}

		id, _ := bucket.NextSequence()
		key := itob(int(id))

		err = bucket.Put(key, value)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

	return err
}

func RetrieveFromDb(bucket []byte, key []byte) error {
	db, err := openDb()
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

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

	return err
}

func RetrieveAllFromDb(bucket []byte, key []byte) error {
	db, err := openDb()
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

	return err
}

func DeleteBucket(bucket []byte) error {
	return nil
}
