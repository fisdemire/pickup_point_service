package httpapi

import (
	"time"

	"github.com/google/uuid"

	"github.com/fisdemire/pickup_point_service/internal/domain"
)

type dummyLoginRequest struct {
	Role string `json:"role"`
}
type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type createPVZRequest struct {
	ID               *uuid.UUID `json:"id"`
	RegistrationDate *time.Time `json:"registrationDate"`
	City             string     `json:"city"`
}
type createReceptionRequest struct {
	PvzID uuid.UUID `json:"pvzId"`
}
type addProductRequest struct {
	Type  string    `json:"type"`
	PvzID uuid.UUID `json:"pvzId"`
}

type errorDTO struct {
	Message string `json:"message"`
}
type userDTO struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
}
type pvzDTO struct {
	ID               uuid.UUID `json:"id"`
	RegistrationDate time.Time `json:"registrationDate"`
	City             string    `json:"city"`
}
type receptionDTO struct {
	ID       uuid.UUID `json:"id"`
	DateTime time.Time `json:"dateTime"`
	PvzID    uuid.UUID `json:"pvzId"`
	Status   string    `json:"status"`
}
type productDTO struct {
	ID          uuid.UUID `json:"id"`
	DateTime    time.Time `json:"dateTime"`
	Type        string    `json:"type"`
	ReceptionID uuid.UUID `json:"receptionId"`
}
type receptionWithProductsDTO struct {
	Reception receptionDTO `json:"reception"`
	Products  []productDTO `json:"products"`
}
type pvzListItemDTO struct {
	PVZ        pvzDTO                     `json:"pvz"`
	Receptions []receptionWithProductsDTO `json:"receptions"`
}

func toUserDTO(u domain.User) userDTO {
	return userDTO{ID: u.ID, Email: u.Email, Role: string(u.Role)}
}
func toPVZDTO(p domain.PVZ) pvzDTO {
	return pvzDTO{ID: p.ID, RegistrationDate: p.RegistrationDate.UTC(), City: p.City}
}
func toReceptionDTO(r domain.Reception) receptionDTO {
	return receptionDTO{ID: r.ID, DateTime: r.DateTime.UTC(), PvzID: r.PVZID, Status: string(r.Status)}
}
func toProductDTO(p domain.Product) productDTO {
	return productDTO{ID: p.ID, DateTime: p.DateTime.UTC(), Type: p.Type, ReceptionID: p.ReceptionID}
}
func toListDTO(in []domain.PVZWithReceptions) []pvzListItemDTO {
	out := make([]pvzListItemDTO, 0, len(in))
	for _, it := range in {
		recs := make([]receptionWithProductsDTO, 0, len(it.Receptions))
		for _, r := range it.Receptions {
			prods := make([]productDTO, 0, len(r.Products))
			for _, p := range r.Products {
				prods = append(prods, toProductDTO(p))
			}
			recs = append(recs, receptionWithProductsDTO{Reception: toReceptionDTO(r.Reception), Products: prods})
		}
		out = append(out, pvzListItemDTO{PVZ: toPVZDTO(it.PVZ), Receptions: recs})
	}
	return out
}
