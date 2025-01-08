package data

import (
	"time"
)

type Book struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Pubished  int
	Pages     int
	Genres    []string
	Rating    float32
	Version   int32
}
