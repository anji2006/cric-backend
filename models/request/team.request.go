package request

type Team struct {
	Name    string `json:"name" binding:"required" bson:"name"`
	Captain string `json:"captain" bson:"captain" binding:"required"`
}
