package controllers

import (
	"encoding/json"
	"net/http"

	usecases "github.com/gabrielsc98/go-uow/internal/application/use-cases"
)

type CreateCompanyInputDTO struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}

type CreateCompanyController struct {
	createCompany *usecases.CreateCompany
}

func NewCreateCompanyController(createCompany *usecases.CreateCompany) *CreateCompanyController {
	return &CreateCompanyController{createCompany: createCompany}
}

func (c *CreateCompanyController) Handle(w http.ResponseWriter, r *http.Request) {
	var input CreateCompanyInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.createCompany.Execute(usecases.CreateCompanyInput{
		Name:      input.Name,
		Email:     input.Email,
		UserName:  input.UserName,
		UserEmail: input.UserEmail,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
