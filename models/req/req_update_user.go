package req

type ReqUpdateUser struct {
	Name string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email"`
}
