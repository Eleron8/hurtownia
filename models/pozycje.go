package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Pozycje is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Pozycje struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Ilosc        int       `json:"ilosc" db:"ilosc"`
	Rabat        float32   `json:"rabat" db:"rabat"`
	ProduktID    uuid.UUID `json:"produktId" db:"produktId"`
	ZamowienieID uuid.UUID `json:"zamowienieId" db:"zamowienieId"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Pozycje) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Pozycjes is not required by pop and may be deleted
type Pozycjes []Pozycje

// String is not required by pop and may be deleted
func (p Pozycjes) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Pozycje) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Pozycje) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Pozycje) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
