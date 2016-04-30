package muxCustomDB

import (
	"fmt"
	"time"
)

var currentId int
var highFives HighFives

// Initialize Data
func init() {
	RepoCreateHighFive(HighFive{Sender: "Ashley", Receiver: "Fred", Message: "Thank you"})
	RepoCreateHighFive(HighFive{Sender: "George", Receiver: "Mason", Message: "You were helpful"})
}

// RepoFindHighFive searches by index for the HighFive to return.
func RepoFindHighFive(id int) HighFive {
	for _, h := range highFives {
		if h.Id == id {
			return h
		}
	}
	// return empty HighFive if not found
	return HighFive{}
}

// RepoCreateHighFive creates a new HighFive.
func RepoCreateHighFive(h HighFive) HighFive {
	currentId += 1
	h.Id = currentId
	h.Date = time.Now()
	highFives = append(highFives, h)
	return h
}

// RepoDestroyHighFive deletes a HighFive by id.
func RepoDestroyHighFive(id int) error {
	for i, h := range highFives {
		if h.Id == id {
			highFives = append(highFives[:i], highFives[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find High Five with id of %d to delete", id)
}
