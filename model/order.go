package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID     uint64 `json:"order_id"`
	CustomerID  uuid.UUID
	LineItems   []LineItem
	CreatedAt   *time.Time
	ShippedAt   *time.Time
	CompletedAt *time.Time
}

type LineItem struct {
	ItemID   uuid.UUID
	Quantity uint
	Price    uint
}
