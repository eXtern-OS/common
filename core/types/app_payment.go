package types

type Payment int

type PaymentType struct {
	Type  Payment `json:"payment_type"` // Type of payment
	Price float64 `json:"price"`        // obvious
}

const (
	Free Payment = iota
	Paid
	Subscription
)
