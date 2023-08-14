package basemodel

type ResponseRequestPagination struct {
	Message string      `json:"message"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Rows    int         `json:"rows"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Message string `json:"message"`
}
