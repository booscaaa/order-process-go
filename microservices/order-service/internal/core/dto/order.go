package dto

type OrderInput struct {
	ProductID string  `json:"product_id"`
	QTT       float64 `json:"qtt"`
}

type OrderQueueOutput struct {
	ProductID string  `json:"product_id"`
	QTT       float64 `json:"qtt"`
}
