package interfaces

type Repository interface {
	Insert(entity Entity) (Entity, error)
	FindByID(id string) (Entity, error)
	FindAll() ([]Entity, error)
}
