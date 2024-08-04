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
	TotalCount int64   `json:"totalCount"`
	Users      []*User `json:"users"`
}

type User struct {
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	PostCount int64  `json:"postCount"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
