package usecases

import (
	"github.com/gabrielsc98/go-uow/internal/domain/entities"
	"github.com/gabrielsc98/go-uow/internal/infra/repositories"
	"github.com/gabrielsc98/go-uow/internal/infra/uow"
)

type CreateCompany struct {
	uow *uow.UnitOfWork
}

func NewCreateCompany(uow *uow.UnitOfWork) *CreateCompany {
	return &CreateCompany{uow: uow}
}

type CreateCompanyInput struct {
	Name      string
	Email     string
	UserName  string
	UserEmail string
}

func (cc *CreateCompany) Execute(input CreateCompanyInput) error {
	companyRepo := cc.uow.GetRepository("companies").(*repositories.CompanyRepository)
	userRepo := cc.uow.GetRepository("users").(*repositories.UserRepository)

	company := entities.NewCompany(entities.CompanyProps{
		Name:  input.Name,
		Email: input.Email,
	})

	cc.uow.StartTransaction()

	createdCompany, err := companyRepo.Insert(company)
	if err != nil {
		cc.uow.Rollback()
		return err
	}

	user := entities.NewUser(entities.UserProps{
		Name:      input.UserName,
		Email:     input.UserEmail,
		CompanyID: createdCompany.ID,
	})

	_, err = userRepo.Insert(user)
	if err != nil {
		cc.uow.Rollback()
		return err
	}

	cc.uow.Commit()
	return nil
}
