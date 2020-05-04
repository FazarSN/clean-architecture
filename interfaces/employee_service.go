package interfaces

import "github.com/clean-architecture/entity"

type EmployeeServiceInterface interface {
	ListEmployee() ([]entity.Employee, error)
	Select(id int) (entity.Employee, error)
	Create(employee entity.Employee) (entity.Employee, error)
	Update(id int, employee entity.Employee) (entity.Employee, error)
	Delete(id int) error
}
