package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/clean-architecture/entity"
	"github.com/clean-architecture/interfaces"
	"github.com/clean-architecture/utility/response"
	"github.com/julienschmidt/httprouter"
)

type EmployeeHandler struct {
	Service interfaces.EmployeeServiceInterface
}

func NewEmployeeHandler(service interfaces.EmployeeServiceInterface) *EmployeeHandler {
	return &EmployeeHandler{
		Service: service,
	}
}

func (e *EmployeeHandler) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	employees, err := e.Service.ListEmployee()
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteSuccess(w, employees, "")
}

func (e *EmployeeHandler) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	Id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return response.WriteError(w, err)
	}

	employee, err := e.Service.Select(Id)
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteSuccess(w, employee, "")
}

func (e *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return response.WriteError(w, err)
	}

	var employee entity.Employee
	err = json.Unmarshal(b, &employee)
	if err != nil {
		return response.WriteError(w, err)
	}

	employeeNew, err := e.Service.Create(employee)
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteSuccess(w, employeeNew, "")
}

func (e *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	Id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return response.WriteError(w, err)
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return response.WriteError(w, err)
	}

	var employee entity.Employee
	err = json.Unmarshal(b, &employee)
	if err != nil {
		return response.WriteError(w, err)
	}

	employeeNew, err := e.Service.Update(Id, employee)
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteSuccess(w, employeeNew, "")
}

func (e *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	Id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return response.WriteError(w, err)
	}

	err = e.Service.Delete(Id)
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteSuccess(w, nil, "success delete")
}
