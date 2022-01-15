package forms

type PassWordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"` //这里的这个mobile是main函数里定义的mobile
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
