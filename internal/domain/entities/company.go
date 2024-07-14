package entities

type Company struct {
	ID    int
	Name  string
	Email string
}

type CompanyProps struct {
	ID    int
	Name  string
	Email string
}

func NewCompany(props CompanyProps) *Company {
	return &Company{
		ID:    props.ID,
		Name:  props.Name,
		Email: props.Email,
	}
}
