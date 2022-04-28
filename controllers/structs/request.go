package structs

type GetPersonByIDRequest struct {
	IdPerson int `json:"id_person" validate:"not-empty, min=0, callback=checkPerson"`
}
