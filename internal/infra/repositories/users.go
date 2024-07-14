package repositories

import (
	"github.com/gabrielsc98/go-uow/internal/domain/entities"
	"github.com/gabrielsc98/go-uow/internal/infra/db"
)

type UserRepository struct {
	db *db.DB
}

func NewUserRepository(db *db.DB) *UserRepository {
	_, err := db.Query("CREATE TABLE IF NOT EXISTS users (id INT PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), company_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE);")
	if err != nil {
		panic(err)
	}
	return &UserRepository{
		db: db,
	}
}

func (cr *UserRepository) Insert(entity *entities.User) (*entities.User, error) {
	rows, err := cr.db.Query("INSERT INTO users (id, name, email, company_id) VALUES (?, ?, ?, ?);", entity.ID, entity.Name, entity.Email, entity.CompanyID)
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

func (cr *UserRepository) FindByID(id int) (*entities.User, error) {
	rows, err := cr.db.Query("SELECT id, name, email, company_id FROM users WHERE id = ?;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		var id, company_id int
		var name, email string
		err = rows.Scan(&id, &name, &email, &company_id)
		if err != nil {
			return nil, err
		}
		return entities.NewUser(entities.UserProps{
			ID:        id,
			Name:      name,
			Email:     email,
			CompanyID: company_id,
		}), nil
	}
	return nil, nil
}

func (cr *UserRepository) FindAll() ([]*entities.User, error) {
	rows, err := cr.db.Query("SELECT id, name, email, company_id FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*entities.User, 0)
	for rows.Next() {
		var id, company_id int
		var name, email string
		err = rows.Scan(&id, &name, &email, &company_id)
		if err != nil {
			return nil, err
		}
		users = append(users, entities.NewUser(entities.UserProps{
			ID:        id,
			Name:      name,
			Email:     email,
			CompanyID: company_id,
		}))
	}
	return users, nil
}
