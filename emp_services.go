package services

import (
	"GolangmMicroservices/webapp/dao"
	"GolangmMicroservices/webapp/errors"
	"GolangmMicroservices/webapp/model"
)

func GetEmployeeid(empId int64) (*model.Employee, *errors.AppError) {
	return dao.GetEmployee(empId)

}
