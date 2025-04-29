package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Approval struct {
	ID        uuid.UUID       `json:"id" db:"id"`
	FlowID    uuid.UUID       `json:"flow_id" db:"flow_id"`
	FlowName  string          `json:"flow_name" db:"flow_name"`
	Status    string          `json:"status" db:"status"`
	CreatedBy uuid.UUID       `json:"created_by" db:"created_by"`
	UpdatedBy *uuid.UUID      `json:"updated_by,omitempty" db:"updated_by"`
	Comments  json.RawMessage `json:"comments" db:"comments"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at"`
}

// Value implements the driver.Valuer interface for Comments
func (a Approval) Value() (driver.Value, error) {
	return json.Marshal(a.Comments)
}

// Scan implements the sql.Scanner interface for Comments
func (a *Approval) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	a.Comments = json.RawMessage(b) // Store raw JSON
	return nil
}
