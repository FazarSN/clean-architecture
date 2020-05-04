package interfaces

import "github.com/clean-architecture/entity"

type EmployeeRepositoryInterface interface {
	SelectAll() ([]entity.Employee, error)
	Select(id int) (entity.Employee, error)
	Insert(employee entity.Employee) (entity.Employee, error)
	Update(id int, employee entity.Employee) (entity.Employee, error)
	Delete(id int) error
}
