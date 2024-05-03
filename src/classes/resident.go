package classes

import "time"

type Resident struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Telephone string `json:"telephone"`
	Apto int64 `json:"apto"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"update"`
}