package repository

//สร้าง  table
type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      int    `db:"status"`
}

//สร้างเต้าเสียบ(เปรียบเทียบ) เพื่อให้เรียกใช้ ข้อมูลใน table ได้ **เป็นชนิดinterface
type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}
