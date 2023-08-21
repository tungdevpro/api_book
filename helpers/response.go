package helpers

type Response struct {
	StatusCode int         `bson:"code" json:"code"`
	Message    string      `bson:"message" json:"message"`
	Data       interface{} `bson:"data" json:"data"`
}

type ErrorResponse struct {
	StatusCode int    `bson:"code" json:"code"`
	Message    string `bson:"message" json:"message"`
}
