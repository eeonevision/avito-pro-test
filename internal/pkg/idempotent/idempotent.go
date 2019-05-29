/*
Copyright 2019 Vladislav Dmitriyev.
*/

package idempotent

import "sync"

// DB struct keep map as simple data store.
type DB struct {
	sync.Mutex
	data map[int]string
}

// NewDB returns new initialised DB struct.
func NewDB() *DB {
	return &DB{data: map[int]string{}}
}

// Insert returns the ID of inserted string value in DB.
// Mutex defends from concurrent changing of data map.
func (i *DB) Insert(value string) int {
	i.Lock()
	lastID := len(i.data) + 1
	i.data[lastID] = value
	i.Unlock()

	return lastID
}

// GetValueByID return string value by given ID.
func (i *DB) GetValueByID(id int) string {
	return i.data[id]
}
