package db

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
)

//Store represents DAL
type Store struct {
	db *bolt.DB
}

//New opens DB for usage
func New() (*Store, error) {
	db, err := bolt.Open("cian.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("offers"))
		return err
	})
	return &Store{db}, nil
}

//Save persists fetched offer ids
func (s *Store) Save(id int) {
	s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		b.Put(itob(id), itob(1))
		return nil
	})
}

//Exists checks if offer was already persisted
func (s *Store) Exists(id int) bool {
	var found bool
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("offers"))
		res := b.Get(itob(id))
		if res != nil {
			found = true
		}
		return nil
	})
	return found
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//Close releases db handle
func (s *Store) Close() {
	s.db.Close()
}
