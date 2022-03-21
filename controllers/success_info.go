//1120009 - Feliciana Gunadi
package controllers

func successUserInfoProcess() UserResponse {
	var response UserResponse
	response.Status = 200
	response.Message = "Success"
	return response
}
