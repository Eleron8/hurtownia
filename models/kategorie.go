package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Kategorie is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Kategorie struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Nazwa       string    `json:"nazwa" db:"nazwa"`
	Opis        string    `json:"opis" db:"opis"`
	KategoriaID uuid.UUID `json:"kategoriaId" db:"kategoriaId"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (k Kategorie) String() string {
	jk, _ := json.Marshal(k)
	return string(jk)
}

// Kategories is not required by pop and may be deleted
type Kategories []Kategorie

// String is not required by pop and may be deleted
func (k Kategories) String() string {
	jk, _ := json.Marshal(k)
	return string(jk)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (k *Kategorie) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (k *Kategorie) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (k *Kategorie) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
