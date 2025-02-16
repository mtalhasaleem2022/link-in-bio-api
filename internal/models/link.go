package models

import "time"

type Link struct {
	ID        string    `bson:"_id,omitempty"`
	Title     string    `bson:"title"`
	URL       string    `bson:"url"`
	CreatedAt time.Time `bson:"createdAt"`
	ExpiresAt time.Time `bson:"expiresAt"`
	Clicks    int       `bson:"clicks"`
}

type Visit struct {
	LinkID    string    `bson:"linkId"`
	IPAddress string    `bson:"ipAddress"`
	Timestamp time.Time `bson:"timestamp"`
}
