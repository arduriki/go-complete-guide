package product

import "errors"


type Product struct{
	title string
	id string
	price float64
}

func New(title, id string, price float64) (*Product, error) {
	if title == "" || id == "" || price <= 0 {
		return nil, errors.New("title, id and price are required")
	}

	return &Product{
		title: title,
		id: id,
		price: price,
	}, nil
}