package validate

import (
	"api/app/helpers/requests"
	"api/app/models"
	"net/mail"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUserMiddleware[T models.UserType](c *gin.Context) {
	var userParams T
	var errorsArr []string
	requests.GetBodyJson(c.Request.Body, &userParams)
	hasErrorKeys, errorsArr := checkCreateKeys(errorsArr, userParams)
	hasErrorValues, errorsArr := checkCreateValues(errorsArr, userParams)

	if hasErrorKeys || hasErrorValues {
		c.JSON(400, gin.H{
			"errors": errorsArr,
		})
		c.Abort()
	}
	c.Set("UserParams", userParams)
	c.Next()
}

func checkCreateKeys[T models.UserType](errorsArr []string, userParams T) (bool, []string) {

	if userParams.GetName() == "" {
		errorsArr = append(errorsArr, "name is required")
	}

	if userParams.GetPassword() == "" {
		errorsArr = append(errorsArr, "password is required")
	}

	if userParams.GetNickname() == "" {
		errorsArr = append(errorsArr, "nickname is required")
	}

	if userParams.GetEmail() == "" {
		errorsArr = append(errorsArr, "email is required")
	}

	if userParams.GetCellphone() == "" {
		errorsArr = append(errorsArr, "cellphone is required")
	}

	return len(errorsArr) > 0, errorsArr

}

func checkCreateValues[T models.UserType](errorsArr []string, userParams T) (bool, []string) {
	if len(strings.Trim(userParams.GetName(), " ")) < 5 {
		errorsArr = append(errorsArr, "Name must have at least 5 caracters")
	}

	if len(strings.Trim(userParams.GetPassword(), " ")) < 8 {
		errorsArr = append(errorsArr, "Password must have at lest 8 caracters")
	}

	if len(strings.Trim(userParams.GetNickname(), " ")) < 5 {
		errorsArr = append(errorsArr, "Nickname must have at least 5 caracters")
	}

	if !validEmail(userParams.GetEmail()) {
		errorsArr = append(errorsArr, "Email not valid")
	}

	if strings.Trim(userParams.GetCellphone(), " ") == "" {
		errorsArr = append(errorsArr, "Cellphone invalid")
	}

	if userParams.GetId() > 0 {
		errorsArr = append(errorsArr, "To update operations you must use PUT method!")
	}

	return len(errorsArr) > 0, errorsArr

}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
