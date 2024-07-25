package types

type RegisterUserPayload struct {
	FirstName string `json:"firstName"  validate:"required"`
	LastName  string `json:"lastName"   validate:"required"`
	Email     string `json:"email"      validate:"required,email"`
	Password  string `json:"password"   validate:"required,min=8,max=16"`
}

type LoginUserPayload struct {
	Email    string `json:"email"      validate:"required,email"`
	Password string `json:"password"   validate:"required"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"createdAt"`
}

type CreateProductPayload struct {
	Name        string  `json:"name"        validate:"required"`
	Description string  `json:"description" validate:"required"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"       validate:"required,number"`
	Quantity    int     `json:"quantity"    validate:"required,number"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CreatedAt   string  `json:"createAt"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user User) error
}
type ProductStore interface {
	GetProducts() ([]Product, error)
	CreateProduct(product Product) error
}
