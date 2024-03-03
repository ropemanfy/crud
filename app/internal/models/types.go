package models

type (
	UserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UserResponse struct {
		ID string `json:"id"`
	}

	User struct {
		ID    string
		Name  string
		Email string
	}
)
