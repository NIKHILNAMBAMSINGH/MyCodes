package dao

import (
	"GolangmMicroservices/webapp/errors"
	"GolangmMicroservices/webapp/model"
	"fmt"
	"net/http"
)

var (
	employee = map[int64]*model.Employee{
		100: {EmpID: 101, EmpName: "Deepak", EmpEmail: "deepak@fmail.com"},
	}
)

func GetEmployee(empId int64) (*model.Employee, *errors.AppError) {
	if employee := employee[empId]; employee != nil {
		return employee, nil
	}
	return nil, &errors.AppError{
		Message:    fmt.Sprintf("Employee %v was not found", empId),
		StatusCode: http.StatusNotFound,
		Status:     "not found",
	}

}
