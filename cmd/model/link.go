package model

import "time"

type Link struct {
	ShortCode string `json:"shortCode" bson:"shortCode"`
	Address   string `json:"address" bson:"address"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
