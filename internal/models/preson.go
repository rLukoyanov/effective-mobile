package model

import "time"

type Person struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Surname     string    `json:"surname" validate:"required"`
	Patronymic  string    `json:"patronymic,omitempty"`
	Age         int       `json:"age,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Nationality string    `json:"nationality,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PersonInput struct {
	Name       string `json:"name" validate:"required"`
	Surname    string `json:"surname" validate:"required"`
	Patronymic string `json:"patronymic,omitempty"`
}

type PersonFilter struct {
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	Patronymic  string `json:"patronymic,omitempty"`
	AgeMin      int    `json:"age_min,omitempty"`
	AgeMax      int    `json:"age_max,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	Page        int    `json:"page,omitempty"`
	PageSize    int    `json:"page_size,omitempty"`
}
