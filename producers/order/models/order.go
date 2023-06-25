package models

type Order struct {
	OrderID    int
	CustomerID int
	ItemID     int
	Quantity   int
	Price      float64
	Total      float64
}
