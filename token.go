package main

import (
	"github.com/Member1221/plutobot-go/db"
	"github.com/boltdb/bolt"
	"errors"
	"fmt"
)

func RegisterToken (db db.PlutoDB, token string) error {
	err := db.Database.Update(func (tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("credentials"))
		if err != nil {
			return err
		}
		bucket := tx.Bucket([]byte("credentials"))
		err = bucket.Put([]byte("token"), []byte(token))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func Token(db db.PlutoDB) (string, error) {
	var token []byte
	err := db.Database.Update(func (tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("credentials"))
		if err != nil {
			return err
		}
		bucket := tx.Bucket([]byte("credentials"))
		token = bucket.Get([]byte("token"))
		if token == nil {
			return errors.New("No tokens was found in " + db.Name + "/credentials/token")
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return string(token)
}
