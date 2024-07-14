package repositories

import (
	"github.com/gabrielsc98/go-uow/internal/domain/entities"
	"github.com/gabrielsc98/go-uow/internal/infra/db"
)

type CompanyRepository struct {
	db *db.DB
}

func NewCompanyRepository(db *db.DB) *CompanyRepository {
	_, err := db.Query("CREATE TABLE IF NOT EXISTS companies (id INT PRIMARY KEY, name VARCHAR(255), email VARCHAR(255));")
	if err != nil {
		panic(err)
	}
	return &CompanyRepository{
		db: db,
	}
}

func (cr *CompanyRepository) Insert(entity *entities.Company) (*entities.Company, error) {
	rows, err := cr.db.Query("INSERT INTO companies (name, email) VALUES (?, ?);", entity.ID, entity.Name, entity.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		entity.ID = id
	}
	return entity, nil
}

func (cr *CompanyRepository) FindByID(id int) (*entities.Company, error) {
	rows, err := cr.db.Query("SELECT id, name, email FROM companies WHERE id = ?;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		var id int
		var name, email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			return nil, err
		}
		return entities.NewCompany(entities.CompanyProps{
			ID:    id,
			Name:  name,
			Email: email,
		}), nil
	}
	return nil, nil
}

func (cr *CompanyRepository) FindAll() ([]*entities.Company, error) {
	rows, err := cr.db.Query("SELECT id, name, email FROM companies;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	companies := make([]*entities.Company, 0)
	for rows.Next() {
		var id int
		var name, email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			return nil, err
		}
		companies = append(companies, entities.NewCompany(entities.CompanyProps{
			ID:    id,
			Name:  name,
			Email: email,
		}))
	}
	return companies, nil
}
