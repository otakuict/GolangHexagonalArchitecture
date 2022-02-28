package service

import (
	"database/sql"
	"errors"
	"hexa/repository"
	"log"
)

//สร้าง struct ของฝั่ง service ซึ่งในทีนี้จะอ้าง struct เดิม ของฝั่ง repo แต่ว่าอ้างได้แค่ interface ที่ทำไว้เท่านั้น
type customerService struct {
	custRepo repository.CustomerRepository
}

//สร้าง fucntion เรียกใช้ struc ด้านบน (ล้อตามฝั่งrepo)
func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	customers, err := s.custRepo.GetAll()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}

		log.Println(err)
		return nil, err
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}

		log.Println(err)
		return nil, err
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
