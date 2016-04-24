// Package high5 provides rest APIs for interacting with High 5's.
package mux

import (
	"time"
)

type HighFive struct {
	Id       int       `json:"id"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	Message  string    `json:"message"`
	Date     time.Time `json:"date"`
}

type HighFives []HighFive
