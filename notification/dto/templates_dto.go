package dto

type OrderCreatedTemplate struct {
	Title    string
	Username string
	OrderID  string
	Amount   string
	Currency string
	Items    []OrderItemTemplate
}

type OrderItemTemplate struct {
	Name     string
	Quantity int
	Price    string
}
