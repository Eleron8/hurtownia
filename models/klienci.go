package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Klienci is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Klienci struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Nazwa     string    `json:"nazwa" db:"nazwa"`
	Imie      string    `json:"imie" db:"imie"`
	Nazwisko  string    `json:"nazwisko" db:"nazwisko"`
	NIP       string    `json:"nip" db:"nip"`
	Adres     string    `json:"adres" db:"adres"`
	Email     string    `json:"email" db:"email"`
	Telefon   string    `json:"telefon" db:"telefon"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (k Klienci) String() string {
	jk, _ := json.Marshal(k)
	return string(jk)
}

// Kliencis is not required by pop and may be deleted
type Kliencis []Klienci

// String is not required by pop and may be deleted
func (k Kliencis) String() string {
	jk, _ := json.Marshal(k)
	return string(jk)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (k *Klienci) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (k *Klienci) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (k *Klienci) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
