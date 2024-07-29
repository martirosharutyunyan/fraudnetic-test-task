package dto

type Page[T any] struct {
	Count int `json:"count"`
	Data  []T `json:"data"`
}

func NewPage[T any](data []T, count int) *Page[T] {
	return &Page[T]{
		Count: count,
		Data:  data,
	}
}
