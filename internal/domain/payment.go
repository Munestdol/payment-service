package domain

type PaymentInfo struct {
	OrderId      string  `json:"orderId" db:"order_id" validate:"required,uuid"`
	UserId       int     `json:"userId" db:"user_id" validate:"required"`
	CardNumber   string  `json:"cardNumber"  validate:"required,alphanum,len=16"`
	CVV          string  `json:"cvv" validate:"required,alphanum,len=3"`
	CardName     string  `json:"cardName" validate:"required,alpha"`
	CardLastName string  `json:"cardLastname" validate:"required,alpha"`
	CardDate     string  `json:"cardDate"`
	TotalPrice   float64 `json:"totalPrice" db:"cost" validate:"required"`
	PaymentType  string  `json:"paymentType" db:"paymentType" validate:"required"`
}

type Transaction struct {
	Id         string  `json:"_" db:"id"`
	UserId     int     `json:"_" db:"user_id"`
	OrderID    string  `json:"_" db:"order_id"`
	CardNumber string  `json:"card_number"`
	Status     string  `json:"status" db:"status"`
	TotalPrice float64 `json:"totalPrice" db:"cost"`
	Date       string  `json:"date" db:"date"`
}
