package dto

type AuthField interface {
	IsEmpty() bool
}

type Register struct {
	Email    string `json:"email" form:"email" valid:"email"`
	FullName string `json:"fullname" form:"fullname" valid:"-"`
	Password string `json:"password" form:"password" valid:"-"`
}

func (r *Register) IsEmpty() bool {
	return r.Email == "" || r.Password == ""
}

type Login struct {
	Email    string `json:"email" form:"email" valid:"email"`
	Password string `json:"password" form:"password" valid:"-"`
}


func (r *Login) IsEmpty() bool {
	return r.Email == "" || r.Password == ""
}
