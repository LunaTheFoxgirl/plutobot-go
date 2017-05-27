package db

import (
	"github.com/bwmarrin/discordgo"
	"github.com/boltdb/bolt"
	"github.com/Member1221/plutobot-go/core"
)

const layout = "06-01-02"

func (p PlutoDB) Addmessage(m *discordgo.Message) {
	p.Database.Update(func(tx *bolt.Tx) error {
		// Create a bucket.
		b, err := tx.CreateBucketIfNotExists([]byte("messages"))
		if err != nil {
			return err
		}

		ts, err := m.Timestamp.Parse()
		if err != nil {
			core.LogError("Error in parsing time for message "+m.ID+": "+ err.Error(), "PB_STORE")
			return err
		}
		timeBucket := ts.Format(layout)

		dayBucket, err := b.CreateBucketIfNotExists([]byte(timeBucket))
		if err != nil {
			core.LogError("Unknown error in creating or getting bucket for "+timeBucket+": " + err.Error(), "PB_STORE")
			return err
		}

		B, err := p.EncodeData(m)
		if err != nil {
			core.LogError("Unknown error in encoding message "+m.ID+": " + err.Error(), "PB_STORE")
			return err
		}
		dayBucket.Put([]byte(m.ID), B)
		return nil
	})
}
