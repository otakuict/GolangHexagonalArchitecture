package service

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"psw"`
}

//go:generate mockgen -destination=../mock/mock_service/mock_customer_service.go bank/service CustomerService
type UserService interface {
	GetUser(string) (*UserResponse, error)
}
