package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" binding:"required" bson:"name"`
	Captain string             `json:"captain" bson:"captain" binding:"required"`
}

type TeamAllRes struct {
	Data  []Team `json:"data"`
	Total int    `json:"total"`
}
