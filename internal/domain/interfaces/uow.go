package interfaces

type UnitOfWork interface {
	RegisterRepository(name string, repository Repository) error
	GetRepository(name string) Repository
	StartTransaction() error
	Commit() error
	Rollback() error
}
