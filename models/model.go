package model

type Customer struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     int    `json:"phone,omitempty"`
	Contacted bool   `json:"contacted,omitempty"`
}

var Customers = map[string]Customer{
	"1": {
		ID:        "1",
		Name:      "Ric",
		Role:      "Software Engineer",
		Email:     "ric@example.com",
		Phone:     5550100,
		Contacted: true,
	},
	"2": {
		ID:        "2",
		Name:      "Hunter",
		Role:      "Product Manager",
		Email:     "hunter@example.com",
		Phone:     5550101,
		Contacted: true,
	},
	"3": {
		ID:        "3",
		Name:      "David",
		Role:      "Director of Sales",
		Email:     "david@example.com",
		Phone:     5550102,
		Contacted: false,
	},
	"4": {
		ID:        "4",
		Name:      "Randy",
		Role:      "People Operatins Manager",
		Email:     "randy@example.com",
		Phone:     5550103,
		Contacted: false,
	},
}
