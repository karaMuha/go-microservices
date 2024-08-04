package models

import "time"

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Category  string    `bson:"category" json:"category"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
