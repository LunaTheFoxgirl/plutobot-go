package db

import (
	"errors"
	"github.com/Member1221/plutobot-go/core"
	"github.com/boltdb/bolt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
)

const layout = "06-01-02"

func (p PlutoDB) AddMessage(m *discordgo.Message) {
	p.Database.Update(func(tx *bolt.Tx) error {
		// Create a bucket.
		b, err := tx.CreateBucketIfNotExists([]byte("messages"))
		if err != nil {
			return err
		}

		ts, err := m.Timestamp.Parse()
		if err != nil {
			core.LogError("Error in parsing time for message "+m.ID+": "+err.Error(), "PB_STORE")
			return err
		}
		timeBucket := ts.Format(layout)

		dayBucket, err := b.CreateBucketIfNotExists([]byte(timeBucket))
		if err != nil {
			core.LogError("Unknown error in creating or getting bucket for "+timeBucket+": "+err.Error(), "PB_STORE")
			return err
		}

		M := Message{m, nil, false, nil}

		B, err := p.EncodeData(M)
		if err != nil {
			core.LogError("Unknown error in encoding message "+m.ID+": "+err.Error(), "PB_STORE")
			return err
		}
		dayBucket.Put([]byte(m.ID), B)
		return nil
	})
}

var (
	ERR_DATEZERO             = errors.New("Error, date is zero.")
	ERR_NOTEXIST_MESSAGEROOT = errors.New("Error, ROOT messages bucket does not exist")
	ERR_NOTEXIST_QUERYBUCKET = errors.New("Error, queried bucket does not exist")
	ERR_NONEXIST_MESSAGEBYID = errors.New("Error, the given ID does not exist within the bucket it should inhabit")
	ERR_PARSE_TIMEFROMID     = func(id string, err error) error {
		return errors.New("Error parsing id " + id + "to string:" + err.Error())
	}
	ERR_DECODE_MESSAGE = func(id string, err error) error {
		return errors.New("Error parsing message " + id + "to message struct: " + err.Error())
	}
)

func (p PlutoDB) IdDate(id string) time.Time {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return time.Time{}
	}
	return time.Unix(((i/4194304)+int64(1420070400000))/1000, 0)
}

func (p PlutoDB) GetMessage(id string) (*Message, error) {
	var m = &Message{}
	err := p.Database.View(func(tx *bolt.Tx) error {
		date := p.IdDate(id)
		if date.IsZero() {
			return ERR_DATEZERO
		}

		timeBucket := date.Format(layout)

		b := tx.Bucket([]byte("messages"))
		if b != nil {
			return ERR_NOTEXIST_MESSAGEROOT
		}

		tb := b.Bucket([]byte(timeBucket))
		if tb != nil {
			return ERR_NOTEXIST_QUERYBUCKET
		}

		k, v := tb.Cursor().Seek([]byte(id))
		if k == nil || string(k) != id {
			return ERR_NONEXIST_MESSAGEBYID
		}

		p.DecodeData(v, m)

		return nil
	})

	return m, err
}

func (p PlutoDB) GetMessages(from, till time.Time) ([]*Message, error) {
	var ms []*Message
	err := p.Database.View(func(tx *bolt.Tx) error {
		mr := tx.Bucket([]byte("messages"))
		if mr != nil {
			return ERR_NOTEXIST_MESSAGEROOT
		}

		for t := time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, nil); t.Before(till); t = t.AddDate(0, 0, 1) { // ranges over the dates
			timeBucket := t.Format(layout)

			tb := mr.Bucket([]byte(timeBucket))
			if tb != nil {
				continue
			}

			c := tb.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				at, err := time.Parse(layout, string(k))
				if err != nil {
					return ERR_PARSE_TIMEFROMID(string(k), err)
				}

				if at.Before(from) || at.After(till) {
					continue
				}

				m := &Message{}
				err = p.DecodeData(v, m)
				if err != nil {
					return ERR_DECODE_MESSAGE(string(k), err)
				}
				ms = append(ms, m)
			}
		}

		return nil
	})

	return ms, err
}

type Message struct {
	M           *discordgo.Message
	Edits       []*MessageEdit
	Deleted     bool
	DeletedWhen time.Time // nil if not caught
}

type MessageEdit struct {
	At         time.Time
	NewContent string
}
