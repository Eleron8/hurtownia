package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Sprzedawcy is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Sprzedawcy struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Nazwa     string    `json:"nazwa" db:"nazwa"`
	NIP       string    `json:"nip" db:"nip"`
	Adres     string    `json:"adres" db:"adres"`
	Email     string    `json:"email" db:"email"`
	Telefon   string    `json:"telefon" db:"telefon"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (s Sprzedawcy) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Sprzedawcies is not required by pop and may be deleted
type Sprzedawcies []Sprzedawcy

// String is not required by pop and may be deleted
func (s Sprzedawcies) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Sprzedawcy) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Sprzedawcy) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Sprzedawcy) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
