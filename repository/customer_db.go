package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB //สร้าง struct database ด้วย pointer
}

//-----------สร้าง  function เพื่อให้คนเรียกใช้inputตัวdatabase  และreturn ออกเป็น customerRepositoryDB struct---------
//เนื่องจากไม่ต้องการให้คนอื่นสามารถมองเห็น struct DB ได้โดยตรง
func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

//-------------------------------------------------------------

//มีเต้าเสียบต้องมีปลั๊กที่สามารถเสียบรูนั้นได้ ต้องทำการ implementหรือconform ตัวinterface ตัวนั้นๆก่อน ในที่นี้คือ  GetAll และ GetByid ของ  CustomerRepository interface

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id=?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
