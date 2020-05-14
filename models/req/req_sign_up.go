package req

type ReqSignUp struct {
	Name string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"required,eqfield=Password"`
}
