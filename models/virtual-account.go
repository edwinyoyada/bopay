package models

import "time"

type VirtualAccount struct {
	ID             string     `json:"id,omitempty"` // will be empty when requesting
	BankCode       string     `json:"bank_code"`
	IsClosed       bool       `json:"is_closed,omitempty"`       // might be empty when requesting
	ExpectedAmount int32      `json:"expected_amount,omitempty"` // might be empty when requesting
	ExternalID     string     `json:"external_id"`
	AccountNumber  string     `json:"account_number,omitempty"` // might be empty when requesting
	Name           string     `json:"name"`
	ExpirationDate *time.Time `json:"expiration_date,omitempty"` // might be empty when requesting
	Status         string     `json:"status,omitempty"`          // might be empty when requesting
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}
