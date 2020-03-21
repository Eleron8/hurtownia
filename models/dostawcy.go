package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Dostawcy is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Dostawcy struct {
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
func (d Dostawcy) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Dostawcies is not required by pop and may be deleted
type Dostawcies []Dostawcy

// String is not required by pop and may be deleted
func (d Dostawcies) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Dostawcy) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Dostawcy) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Dostawcy) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
