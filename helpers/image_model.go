package helpers

type Image struct {
	Id        string `bson:"id" json:"id"`
	Url       string `bson:"url" json:"url"`
	Width     int    `bson:"width" json:"width"`
	Height    int    `bson:"height" json:"height"`
	Extension string `bson:"extension" json:"extension"`
}
