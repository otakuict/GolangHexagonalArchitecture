package repository

//สร้าง  table
type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

//สร้างเต้าเสียบ(เปรียบเทียบ) เพื่อให้เรียกใช้ ข้อมูลใน table ได้ **เป็นชนิดinterface
type UserRepository interface {
	GetUser(string) (*User, error)
}
