package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Guild struct {
	Id primitive.ObjectID `json:"id,omitempty"`

	// TODO: Finish the guild data for base data information and configuration for scales
}