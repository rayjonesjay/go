package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record // slice of type Record
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"` // location from index 0
}

// NewLog returns a pointer to a new empty Log
func NewLog() *Log {
	return &Log{}
}

// Method: Append adds a Record to Log.records slice
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// Method: Read attempts to read a record at offset, if offset is greater than or
// equal to length of records, Read returns an empty Record and ErrOffsetNotFound error
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
