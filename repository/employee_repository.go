package repository

import (
	"database/sql"
	"net/http"

	"github.com/clean-architecture/entity"
	customErr "github.com/clean-architecture/utility/error"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (e *EmployeeRepository) SelectAll() ([]entity.Employee, error) {
	rows, err := e.db.Query("SELECT id, firstname, lastname, city, street, company_id FROM employees")
	defer rows.Close()

	if err != nil {
		return []entity.Employee{}, customErr.New(err.Error(), http.StatusInternalServerError)
	}

	employee := entity.Employee{}
	employees := []entity.Employee{}

	for rows.Next() {
		err = rows.Scan(
			&employee.ID,
			&employee.Firstname,
			&employee.Lastname,
			&employee.City,
			&employee.Street,
			&employee.CompanyID,
		)
		if err != nil {
			return []entity.Employee{}, customErr.New(err.Error(), http.StatusInternalServerError)
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (e *EmployeeRepository) Select(id int) (entity.Employee, error) {
	rows, err := e.db.Query("SELECT id, firstname, lastname, city, street, company_id FROM employees WHERE id = ?", id)
	defer rows.Close()

	if err != nil {
		return entity.Employee{}, customErr.New(err.Error(), http.StatusInternalServerError)
	}

	employee := entity.Employee{}

	isExist := rows.Next()
	if !isExist {
		return entity.Employee{}, customErr.New("not found", http.StatusNotFound)
	}

	err = rows.Scan(
		&employee.ID,
		&employee.Firstname,
		&employee.Lastname,
		&employee.City,
		&employee.Street,
		&employee.CompanyID,
	)
	if err != nil {
		return entity.Employee{}, customErr.New(err.Error(), http.StatusInternalServerError)
	}

	return employee, nil
}

func (e *EmployeeRepository) Insert(employee entity.Employee) (entity.Employee, error) {
	result, err := e.db.Exec("INSERT INTO employees (firstname, lastname, city, street, company_id) VALUES (?, ?, ?, ?, ?)", employee.Firstname, employee.Lastname, employee.City, employee.Street, employee.CompanyID)
	if err != nil {
		return entity.Employee{}, customErr.New(err.Error(), http.StatusInternalServerError)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.Employee{}, customErr.New(err.Error(), http.StatusInternalServerError)
	}

	employee.ID = int(id)

	return employee, nil
}

func (e *EmployeeRepository) Update(id int, employee entity.Employee) (entity.Employee, error) {
	_, err := e.db.Exec("UPDATE employees SET firstname = ?, lastname = ?, city = ?, street = ?, company_id = ? WHERE id = ?", employee.Firstname, employee.Lastname, employee.City, employee.Street, employee.CompanyID, id)
	if err != nil {
		return entity.Employee{}, customErr.New(err.Error(), http.StatusInternalServerError)
	}

	return employee, nil
}

func (e *EmployeeRepository) Delete(id int) error {
	_, err := e.db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		return customErr.New(err.Error(), http.StatusInternalServerError)
	}

	return nil
}
