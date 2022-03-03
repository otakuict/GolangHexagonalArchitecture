package service

import (
	"database/sql"
	"errors"
	"hexa/repository"
	"log"
)

//สร้าง struct ของฝั่ง service ซึ่งในทีนี้จะอ้าง struct เดิม ของฝั่ง repo แต่ว่าอ้างได้แค่ interface ที่ทำไว้เท่านั้น
type userService struct {
	userRepo repository.UserRepository
}

//สร้าง fucntion เรียกใช้ struc ด้านบน (ล้อตามฝั่งrepo)
func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) GetUser(username string) (*UserResponse, error) {
	user, err := s.userRepo.GetUser(username)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}

		log.Println(err)
		return nil, err
	}

	userResponse := UserResponse{
		Username: user.Username,
		Password: user.Password,
	}

	return &userResponse, nil
}
