package entity

// Employee represents an user/employee in the system
type User struct {
	BaseModel
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Password   string  `json:"-"` // Hashed password, not exposed in JSON
	BaseSalary float64 `json:"base_salary"`
	IsAdmin    bool    `json:"is_admin"`
}
