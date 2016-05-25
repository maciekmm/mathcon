package models

import "google.golang.org/cloud/datastore"

type ContestType string

const (
	ContestTypeSelect ContestType = "select"
	ContestTypeInput  ContestType = "input"
)

type Contest struct {
	Owner     string
	Name      string
	Type      ContestType
	Questions []*datastore.Key
}
