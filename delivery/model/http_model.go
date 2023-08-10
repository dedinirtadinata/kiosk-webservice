package model

type ResponseSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
}
