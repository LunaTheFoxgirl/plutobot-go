package db

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
)

func Open(name string) (PlutoDB, error) {
	var db PlutoDB = PlutoDB{}
	bdb, err := bolt.Open(name+".pldb", 0600, nil)
	if err != nil {
		return db, err
	}
	db.Database = bdb
	db.Name = name

	return db, nil
}

type PlutoDB struct {
	Database *bolt.DB
	Name     string
}

func (db PlutoDB) EncodeData(data interface{}) ([]byte, error) {
	var output bytes.Buffer
	enc := gob.NewEncoder(&output)
	err := enc.Encode(data)
	if err != nil {
		return []byte{}, err
	}
	return output.Bytes(), nil
}

func (db PlutoDB) DecodeData(data []byte, output interface{}) error {
	var input = bytes.NewBuffer(data)

	enc := gob.NewDecoder(input)

	//Decode the input data, and via the pointer set it at its destination.
	err := enc.Decode(&output)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (db PlutoDB) Close() error {
	err := db.Database.Close()
	if err != nil {
		return err
	}
	return nil
}
