package domain

type PaymentInfo struct {
	OrderId      string  `json:"order_id" db:"order_id" validate:"required,uuid"`
	UserId       string  `json:"user_id" db:"user_id" validate:"required,uuid"`
	CardNumber   string  `json:"card_number"  validate:"required,alphanum,len=16"`
	CVV          string  `json:"cvv" validate:"required,alphanum,len=3"`
	CardName     string  `json:"card_name" validate:"required,alpha"`
	CardLastName string  `json:"card_last_name" validate:"required,alpha"`
	CardDate     string  `json:"card_date"`
	TotalPrice   float64 `json:"total_price" db:"cost" validate:"required"`
}

type Transaction struct {
	Id         string  `json:"_" db:"id"`
	UserId     string  `json:"_" db:"user_id"`
	OrderID    string  `json:"_" db:"order_id"`
	CardNumber string  `json:"card_number"`
	Status     string  `json:"status" db:"status"`
	TotalPrice float64 `json:"total_price" db:"cost"`
	Date       string  `json:"date" db:"date"`
}
