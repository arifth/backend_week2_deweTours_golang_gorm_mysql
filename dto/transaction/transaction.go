package transactiondto

type CreateTransactionRequest struct {
	// Name     string `json:"name" form:"name" validate:"required"`
	// Email    string `json:"email" form:"email" validate:"required"`
	// Password string `json:"password" form:"password" validate:"required"`

	CounterQty int    `json:"counter_qty"`
	Total      int    `json:"total"`
	Status     string `json:"status"`
	Attachment string `json:"attachment"`
	TripId     int    `json:"trip_id"`
	// Trip       `json:"trip"`
	// UserId     int          `json:"user_id"`
	// User       UserResponse `json:"user"`
}
