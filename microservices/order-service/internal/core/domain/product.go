package domain

import "context"

type Product struct {
	ID          string
	Description string
	Price       float64
}

type ProductDBRepository interface {
	Fetch(context.Context) (*[]Product, error)
}
