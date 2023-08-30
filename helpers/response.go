package helpers

type Response struct {
	StatusCode int         `bson:"code" json:"code"`
	Message    string      `bson:"message" json:"message"`
	Data       interface{} `bson:"data" json:"data"`
	Paging interface{} `bson:"paging" json:"paging,omitempty"`
	Filter interface{} `bson:"filter" json:"filter,omitempty"`
}

func SimpleSuccessRes(data interface{}) *Response {
	return NewSuccessResponse(data, nil, nil)
}

func NewSuccessResponse(data,paging, filter  interface{}) *Response  {
	return &Response{
		StatusCode: 1,
		Message: "Success",
		Data: data,
		Paging: paging,
		Filter: filter,
	}
}

type ErrorResponse struct {
	StatusCode int    `bson:"code" json:"code"`
	Message    string `bson:"message" json:"message"`
}
