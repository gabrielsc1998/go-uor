package uow

import (
	"errors"

	"github.com/gabrielsc98/go-uow/internal/infra/db"
)

type Repository interface{}

type UnitOfWork struct {
	db                 *db.DB
	transactionStarted bool
	repositories       map[string]Repository
}

func NewUnitOfWork(db *db.DB) *UnitOfWork {
	return &UnitOfWork{
		db:           db,
		repositories: make(map[string]Repository),
	}
}

func (uow *UnitOfWork) RegisterRepository(name string, repository Repository) {
	uow.repositories[name] = repository
}

func (uow *UnitOfWork) GetRepository(name string) Repository {
	return uow.repositories[name]
}

func (uow *UnitOfWork) StartTransaction() error {
	if uow.transactionStarted {
		return errors.New("transaction already started")
	}
	uow.transactionStarted = true
	_, err := uow.db.Query("START TRANSACTION;")
	return err
}

func (uow *UnitOfWork) Commit() error {
	_, err := uow.db.Query("COMMIT;")
	if err != nil {
		uow.Rollback()
	}
	uow.transactionStarted = true
	return err
}

func (uow *UnitOfWork) Rollback() error {
	_, err := uow.db.Query("ROLLBACK;")
	if err != nil {
		return err
	}
	uow.transactionStarted = false
	return nil
}
