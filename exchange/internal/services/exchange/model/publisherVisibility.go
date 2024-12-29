package model

type PublisherVisibility struct {
	PublisherID string `bson:"_id"`
	Loads       uint32 `bson:"loads"`
	Views       uint32 `bson:"views"`
}

func (s *PublisherVisibility) GetVisibility() float32 {
	return float32(s.Views) / float32(s.Loads)
}
