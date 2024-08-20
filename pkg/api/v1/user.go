package v1

type RegistUserRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
	Nickname string `json:"nickname" valid:"required,stringlength(1|255)"`
	Email    string `json:"email" valid:"required,email"`
	Phone    string `json:"phone" valid:"required,stringlength(11|11)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
}

type RegistUserResponse User

type LoginUserRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type GetUserResponse User

type ListUserRequest struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type ListUserResponse struct {
	Total int64   `json:"total"`
	Users []*User `json:"users"`
}

type UpdateUserRequest struct {
	Nickname *string `json:"nickname" valid:"stringlength(1|255)"`
	Email    *string `json:"email" valid:"email"`
	Phone    *string `json:"phone" valid:"stringlength(11|11)"`
}

type UpdateUserResponse User

type DeleteUserResponse User

type User struct {
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	RoleID    string `json:"role_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
