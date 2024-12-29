package model

type GeoModel struct {
	Country string `json:"country" bson:"country"`
	Region  string `json:"region" bson:"region"`
	City    string `json:"city" bson:"city"`
}
