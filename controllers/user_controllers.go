//1120009 - Feliciana Gunadi
package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//this function to get all data from a user by user id
func GetAUserData(idUser int) (User, error) {
	db := connect()
	defer db.Close()

	var user User
	rows := db.QueryRow("SELECT * FROM users WHERE id=?", idUser)
	err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Password, &user.Email, &user.UserType)
	if err != nil {
		log.Println("Something wrong when get a user old data")
	}
	return user, err
}

//this function to fill user output info with user real info, because I'm not show user password when show user data
func GetAUserInfo(user User) UserInfo {
	var userInfo UserInfo
	userInfo.Name = user.Name
	userInfo.Age = user.Age
	userInfo.Address = user.Address
	userInfo.ID = user.ID
	userInfo.Email = user.Email
	userInfo.UserType = user.UserType
	return userInfo
}

//this function to check user login activity
func GetUserLogin(c echo.Context) error {
	db := connect()
	defer db.Close()
	// user := new(UserInfo)

	// err := r.ParseForm()
	// if err != nil {
	// 	log.Println(err)
	// 	response := errorParseFormInfo()
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode((response))
	// 	return c.JSON(http.StatusOK, response)
	// }

	password := encodePassword(c.FormValue("password"))
	fmt.Println(password)
	email := c.FormValue("email")

	if password != "" && email != "" {
		query := "SELECT * FROM users WHERE password = ? AND email = ?"

		rows, err := db.Query(query, password, email)
		if err != nil {
			log.Println(err)
			response := errorQueryInfo(0)
			return c.JSON(http.StatusOK, response)
		}

		var user User
		if rows.Next() {
			err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Password, &user.Email, &user.UserType)
		}

		if err != nil {
			log.Println(err.Error())
			response := errorTableField()
			return c.JSON(http.StatusOK, response)
		}

		if user.Name != "" {
			response := successUserInfoProcess()
			response.Data = GetAUserInfo(user)
			generateToken(c, user.ID, user.Name, user.UserType)
			return c.JSON(http.StatusOK, response)
		} else {
			response := errorUserNotFound()
			return c.JSON(http.StatusOK, response)
		}
	} else {
		response := errorEmptyForm()
		return c.JSON(http.StatusOK, response)
	}
}

//this function to get all user data, so it showed array of user
func GetAllUsers(c echo.Context) error {
	db := connect()
	defer db.Close()

	query := "SELECT id, name, age, address, email, userType FROM users"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		response := errorQueryInfo(0)
		return c.JSON(http.StatusOK, response)
	}

	var user UserInfo
	var users []UserInfo
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.UserType); err != nil {
			log.Println(err.Error())
			response := errorTableField()
			return c.JSON(http.StatusOK, response)
		} else {
			users = append(users, user)
		}
	}

	if len(users) != 0 {
		var response UsersResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = users
		return c.JSON(http.StatusOK, response)
	} else {
		response := errorEmptyArray()
		return c.JSON(http.StatusOK, response)
	}
}

//this function to insert a user, and check insert process is success or not
func InsertUser(c echo.Context) error {
	db := connect()
	defer db.Close()

	// err := r.ParseForm()
	// if err != nil {
	// 	log.Println(err)
	// 	response := errorParseFormInfo()
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode((response))
	// 	return c.JSON(http.StatusOK, response)
	// }

	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")
	password := encodePassword(c.FormValue("password"))
	email := c.FormValue("email")
	// userType := c.FormValue("user type")

	if name != "" && age != 0 && address != "" && password != "" && email != "" {
		result, errQuery := db.Exec("INSERT INTO users(name, age, address, password, email) values (?,?,?,?,?)", name, age, address, password, email)
		if errQuery == nil {
			response := successUserInfoProcess()
			userId, _ := result.LastInsertId()
			user, _ := GetAUserData(int(userId))
			response.Data = GetAUserInfo(user)
			return c.JSON(http.StatusOK, response)
		} else {
			response := errorQueryInfo(1)
			return c.JSON(http.StatusOK, response)
		}
	} else {
		response := errorEmptyForm()
		return c.JSON(http.StatusOK, response)
	}
}

//this function to update a user data, and check user update is success or not
func UpdateUser(c echo.Context) error {
	db := connect()
	defer db.Close()

	userId, _ := strconv.Atoi(c.Param("user_id"))
	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")
	password := encodePassword(c.FormValue("password"))
	tempEmptyString := encodePassword("")
	email := c.FormValue("email")
	fmt.Println(userId)
	tempUser, err := GetAUserData(userId)

	if err == nil {
		if name == "" {
			name = tempUser.Name
		}
		if age == 0 {
			age = tempUser.Age
		}
		if address == "" {
			address = tempUser.Address
		}
		if password == tempEmptyString {
			password = tempUser.Password
		}
		if email == "" {
			email = tempUser.Email
		}
		if name == tempUser.Name && age == tempUser.Age && address == tempUser.Address && password == tempUser.Password && email == tempUser.Email {
			response := errorUpdateForm()
			return c.JSON(http.StatusOK, response)
		} else {
			_, errQuery := db.Exec("UPDATE users SET name = ?, age = ?, address = ?, password = ?, email = ? WHERE id = ?", name, age, address, password, email, userId)
			if errQuery == nil {
				response := successUserInfoProcess()
				user, _ := GetAUserData(userId)
				response.Data = GetAUserInfo(user)
				return c.JSON(http.StatusOK, response)
			} else {
				response := errorQueryInfo(2)
				return c.JSON(http.StatusOK, response)
			}
		}
	} else {
		response := errorIDNotFound()
		return c.JSON(http.StatusOK, response)
	}
}

//this function to delete a user from database, and show delete process is success or not
func DeleteUser(c echo.Context) error {
	db := connect()
	defer db.Close()

	userId, _ := strconv.Atoi(c.Param("user_id"))
	user, err := GetAUserData(userId)

	if err != nil {
		response := errorIDNotFound()
		return c.JSON(http.StatusOK, response)
	} else {
		_, errDelete := db.Exec("DELETE FROM transactions WHERE userId = ?", userId)
		if errDelete != nil {
			response := errorDataForeignKey()
			return c.JSON(http.StatusOK, response)
		}
		_, errQuery := db.Exec("DELETE FROM users WHERE id=?", userId)
		if errQuery == nil {
			response := successUserInfoProcess()
			response.Data = GetAUserInfo(user)
			return c.JSON(http.StatusOK, response)
		} else {
			response := errorQueryInfo(3)
			return c.JSON(http.StatusOK, response)
		}
	}
}

//this function to encode user password
func encodePassword(pass string) string {
	encodePass := md5.Sum([]byte(pass))
	return hex.EncodeToString(encodePass[:])
}

func UserLogout(c echo.Context) error {
	resetUserToken(c)

	var response UserResponse
	response.Status = 200
	response.Message = "Success"

	return c.JSON(http.StatusOK, response)
}
