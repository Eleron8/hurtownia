package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Kontum is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Kontum struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Login     string    `json:"login" db:"login"`
	Email     string    `json:"email" db:"email"`
	Haslo     string    `json:"haslo" db:"haslo"`
	Salt      string    `json:"salt" db:"salt"`
	KlientID  int       `json:"klientId" db:"klientId"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (k Kontum) String() string {
	jk, _ := json.Marshal(k)
	return string(jk)
}

// Konta is not required by pop and may be deleted
type Konta []Kontum

// String is not required by pop and may be deleted
func (k Konta) String() string {
	jk, _ := json.Marshal(k)
	return string(jk)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (k *Kontum) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (k *Kontum) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (k *Kontum) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
