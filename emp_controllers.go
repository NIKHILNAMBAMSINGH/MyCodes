package controllers

import (
	"GolangmMicroservices/webapp/errors"
	"GolangmMicroservices/webapp/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetEmployee(response http.ResponseWriter, request *http.Request) {
	empId, err := strconv.ParseInt(request.URL.Query().Get("EmpID"), 10, 64)
	if err != nil {
		apiError := &errors.AppError{
			Message:    "emp_id must be a number",
			StatusCode: http.StatusBadRequest,
			Status:     "bad request",
		}
		jsonValue, _ := json.Marshal(apiError)
		response.WriteHeader(apiError.StatusCode)
		response.Write(jsonValue)
		return
	}
	emp, apiError := services.GetEmployeeid(empId)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		response.WriteHeader(apiError.StatusCode)
		response.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(emp)
	response.Write(jsonValue)

}
