package dto

type Register struct {
	Email    string `json:"email" form:"email" valid:"email"`
	FullName string `json:"fullname" form:"fullname" valid:"-"`
	Password string `json:"password" form:"password" valid:"-"`
}

func (r *Register) IsEmpty() bool {
	return r.Email == ""
}
