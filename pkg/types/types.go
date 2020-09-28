package types

// Money t
type Money int64

// Currency t
type Currency string

// PAN t
type PAN string

// Category t
type Category string

// Status t
type Status string

const (
	// StatusOk v
	StatusOk Status = "OK"

	// StatusFail v
	StatusFail Status = "FAIL"

	// StatusInProgress v
	StatusInProgress Status = "INPROGRESS"
)

const (
	// TJS v
	TJS Currency = "TJS"

	// RUB v
	RUB Currency = "RUB"

	// USD v
	USD Currency = "USD"
)

// Card t
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Payment t
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

// PaymentSource t
type PaymentSource struct {
	Type    string
	Number  PAN
	Balance Money
}
