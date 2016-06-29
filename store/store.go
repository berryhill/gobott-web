package store

import (
	//"fmt"
	"log"
	"encoding/binary"

	//"github.com/boltdb/bolt"
	"gopkg.in/mgo.v2"
)

var url string = "http://127.0.0.1:27017"

func openDb() (*mgo.Session, error) {
	//db, err := bolt.Open("my.db", 0600, nil)
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	return session, err
}

func OpenDb() (*mgo.Session, error) {
	//db, err := bolt.Open("my.db", 0600, nil)
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	return session, err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))

	return b
}

func AddToDb(database []byte, data []byte, ) error {
	//err = db.Update(func(tx *bolt.Tx) error {
	//	bucket, err := tx.CreateBucketIfNotExists(bucket)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	id, _ := bucket.NextSequence()
	//	key := itob(int(id))
	//
	//	err = bucket.Put(key, value)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	return nil
	//})
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err = db.Close(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//return err
	return nil
}

func RetrieveFromDb(bucket []byte, key []byte) error {
	//db, err := openDb()
	//defer db.Close()
	//
	//err = db.View(func(tx *bolt.Tx) error {
	//	bucket := tx.Bucket(bucket)
	//	if bucket == nil {
	//		return fmt.Errorf("Bucket %q not found!", key)
	//	}
	//
	//	val := bucket.Get(key)
	//	fmt.Println(string(val))
	//
	//	return nil
	//})
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err = db.Close(); err != nil {
	//	log.Fatal(err)
	//}
	//
	return nil
}

func RetrieveAllFromDb(model interface{}, bucket []byte) /*(map[string][]byte, error)*/ error {
	//db, err := openDb()
	//defer db.Close()
	//
	//err = db.View(func(tx *bolt.Tx) error {
	//	b := tx.Bucket(bucket)
	//	c := b.Cursor()
	//
	//	//var list []interface{}
	//	for k, v := c.First(); k != nil; k, v = c.Next() {
	//		//retrieved := new(model)
	//		//list = append(list, )
	//
	//		fmt.Printf("key=%s, value=%s\n", k, v)
	//	}
	//
	//	return nil
	//})
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err = db.Close(); err != nil {
	//	log.Fatal(err)
	//}

	return nil
}

func DeleteBucket(bucket []byte) error {
	//TODO implement
	return nil
}
