package request

type User struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=25"`
	LastName  string `json:"last_name" binding:"required,min=1,max=25"`
	Email     string `json:"email" binding:"email"`
	Password  string `json:"password" binding:"min=6,max=25"`
}
