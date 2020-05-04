package service

import (
	"github.com/clean-architecture/entity"
	"github.com/clean-architecture/interfaces"
)

type EmployeeService struct {
	Repo interfaces.EmployeeRepositoryInterface
}

func NewEmployeeService(repo interfaces.EmployeeRepositoryInterface) *EmployeeService {
	return &EmployeeService{
		Repo: repo,
	}
}

func (e *EmployeeService) ListEmployee() ([]entity.Employee, error) {
	return e.Repo.SelectAll()
}

func (e *EmployeeService) Select(id int) (entity.Employee, error) {
	return e.Repo.Select(id)
}

func (e *EmployeeService) Create(employee entity.Employee) (entity.Employee, error) {
	return e.Repo.Insert(employee)
}

func (e *EmployeeService) Update(id int, employee entity.Employee) (entity.Employee, error) {
	return e.Repo.Update(id, employee)
}

func (e *EmployeeService) Delete(id int) error {
	return e.Repo.Delete(id)
}
