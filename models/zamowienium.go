package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Zamowienium is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Zamowienium struct {
	ID             uuid.UUID `json:"id" db:"id"`
	DataZlozenia   time.Time `json:"dataZlozenia" db:"dataZlozenia"`
	DataRealizacji time.Time `json:"dataRealizacji" db:"dataRealizacji"`
	DataWplaty     time.Time `json:"dataWplaty" db:"dataWplaty"`
	KlientID       uuid.UUID `json:"klientId" db:"klientId"`
	SprzedawcaID   uuid.UUID `json:"sprzedawcaId" db:"sprzedawcaId"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (z Zamowienium) String() string {
	jz, _ := json.Marshal(z)
	return string(jz)
}

// Zamowienia is not required by pop and may be deleted
type Zamowienia []Zamowienium

// String is not required by pop and may be deleted
func (z Zamowienia) String() string {
	jz, _ := json.Marshal(z)
	return string(jz)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (z *Zamowienium) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (z *Zamowienium) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (z *Zamowienium) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
