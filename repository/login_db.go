package repository

import "github.com/jmoiron/sqlx"

type UserRepositoryDB struct {
	db *sqlx.DB //สร้าง struct database ด้วย pointer
}

//เนื่องจากไม่ต้องการให้คนอื่นสามารถมองเห็น struct DB ได้โดยตรง
func NewUserRepositoryDB(db *sqlx.DB) UserRepository {
	return UserRepositoryDB{db: db}
}

//-------------------------------------------------------------

func (r UserRepositoryDB) GetUser(username string) (*User, error) {
	user := User{}
	query := "select password from users where username=?"
	err := r.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
