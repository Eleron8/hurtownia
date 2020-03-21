package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Produkty is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Produkty struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Nazwa       string    `json:"nazwa" db:"nazwa"`
	CenaNetto   float32   `json:"cenaNetto" db:"cenaNetto"`
	VAT         string    `json:"vat" db:"vat"`
	Opis        string    `json:"opis" db:"opis"`
	Foto        string    `json:"foto" db:"foto"`
	KategoriaID uuid.UUID `json:"kategoriaId" db:"kategoriaId"`
	DostawcaID  uuid.UUID `json:"dostawcaId" db:"dostawcaId"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Produkty) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Produkties is not required by pop and may be deleted
type Produkties []Produkty

// String is not required by pop and may be deleted
func (p Produkties) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Produkty) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Produkty) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Produkty) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
