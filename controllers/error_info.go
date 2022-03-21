//1120009 - Feliciana Gunadi
package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func errorQueryInfo(num int) ErrorResponse {
	var response ErrorResponse
	response.Status = 400
	if num == 1 {
		response.Message = "Error Query Insert"
	} else if num == 2 {
		response.Message = "Error Query Update"
	} else if num == 3 {
		response.Message = "Error Query Delete"
	} else if num == 0 {
		response.Message = "Error Query Select"
	}
	return response
}

func errorEmptyArray() ErrorResponse {
	var response ErrorResponse
	response.Status = 400
	response.Message = "Error Empty Array"
	return response
}

func errorTableField() ErrorResponse {
	var response ErrorResponse
	response.Status = 400
	response.Message = "Error table field doesn't correct"
	return response
}

func errorEmptyForm() ErrorResponse {
	var response ErrorResponse
	response.Status = 400
	response.Message = "Error Empty Insert Value"
	return response
}

func errorIDNotFound() ErrorResponse {
	var response ErrorResponse
	response.Status = 400
	response.Message = "Error Id Not Found"
	return response
}

func errorUpdateForm() ErrorResponse {
	var response ErrorResponse
	response.Status = 400
	response.Message = "Error Can't Update. Check Your Form"
	return response
}

func errorDataForeignKey() ErrorResponse {
	var response ErrorResponse
	response.Status = 400
	response.Message = "Error Can't Delete Foreign Key"
	return response
}

func errorUserNotFound() ErrorResponse {
	var response ErrorResponse
	response.Status = 400
	response.Message = "User Not Found"
	return response
}

func sendUnAuthorizedResponse(c echo.Context) error {
	var response ErrorResponse
	response.Status = 401
	response.Message = "Unauthorized Access"
	return c.JSON(http.StatusOK, response)
}
