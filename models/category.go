package models

type Category struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CategoryCreateRequest struct {
	Name  string  `json:"name" db:"name"`
	Price float64 `json:"price" db:"price"`
}

type Response struct {
	Message string      `json:message`
	Status  int         `json:status`
	Data    interface{} `json:data`
}
