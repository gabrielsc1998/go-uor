package entities

type User struct {
	ID        int
	Name      string
	Email     string
	CompanyID int
}

type UserProps struct {
	ID        int
	Name      string
	Email     string
	CompanyID int
}

func NewUser(props UserProps) *User {
	return &User{
		ID:        props.ID,
		Name:      props.Name,
		Email:     props.Email,
		CompanyID: props.CompanyID,
	}
}
