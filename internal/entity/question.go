package entity

import "time"

type Answer struct {
	ID    int32
	Text  string
	Media Media
}

type Question struct {
	ID         int32
	Text       string
	Answer     Answer
	Author     string
	Media      Media
	CreateTime time.Time
}