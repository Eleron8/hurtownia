package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Platnosci is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Platnosci struct {
	ID              uuid.UUID `json:"id" db:"id"`
	NrKontaBenef    string    `json:"nrKontaBenef" db:"nrKontaBenef"`
	NrKontaPlatnika string    `json:"nrKontaPlatnika" db:"nrKontaPlatnika"`
	KwotaBrutto     float32   `json:"kwotaBrutto" db:"kwotaBrutto"`
	Waluta          int       `json:"waluta" db:"waluta"`
	DataOtrzymania  time.Time `json:"dataOtrzymania" db:"dataOtrzymania"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Platnosci) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Platnoscis is not required by pop and may be deleted
type Platnoscis []Platnosci

// String is not required by pop and may be deleted
func (p Platnoscis) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Platnosci) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Platnosci) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Platnosci) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
