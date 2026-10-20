package domain

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleEmployee  Role = "employee"
	RoleModerator Role = "moderator"
)

func (r Role) Valid() bool { return r == RoleEmployee || r == RoleModerator }

type User struct {
	ID           uuid.UUID
	Email        string
	Role         Role
	PasswordHash string
}

type PVZ struct {
	ID               uuid.UUID
	RegistrationDate time.Time
	City             string
}

type CreatePVZ struct {
	ID               *uuid.UUID
	RegistrationDate *time.Time
	City             string
}

type ReceptionStatus string

const (
	ReceptionInProgress ReceptionStatus = "in_progress"
	ReceptionClosed     ReceptionStatus = "close"
)

type Reception struct {
	ID       uuid.UUID
	PVZID    uuid.UUID
	DateTime time.Time
	Status   ReceptionStatus
}

type Product struct {
	ID          uuid.UUID
	ReceptionID uuid.UUID
	DateTime    time.Time
	Type        string
}

type ReceptionWithProducts struct {
	Reception Reception
	Products  []Product
}

type PVZWithReceptions struct {
	PVZ        PVZ
	Receptions []ReceptionWithProducts
}

type ListFilter struct {
	From, To    *time.Time
	Page, Limit int
}

var cities = map[string]struct{}{"Москва": {}, "Санкт-Петербург": {}, "Казань": {}}
var productTypes = map[string]struct{}{"электроника": {}, "одежда": {}, "обувь": {}}

func ValidCity(c string) bool        { _, ok := cities[c]; return ok }
func ValidProductType(t string) bool { _, ok := productTypes[t]; return ok }
