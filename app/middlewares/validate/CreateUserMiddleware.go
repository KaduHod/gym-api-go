package validate

import (
	"api/app/helpers/requests"
	"api/app/models"
	"net/mail"
	"strings"

	"github.com/gin-gonic/gin"
)

type CheckUserPostParams[T models.Aluno | models.Personal | models.User] struct {
	UserParams T
}

func CreateUserMiddleware[T models.Aluno | models.Personal | models.User](c *gin.Context) {
	var userParams models.User
	var errorsArr []string
	requests.GetBodyJson(c.Request.Body, &userParams)
	hasErrorKeys, errorsArr := checkKeys(errorsArr, userParams)
	hasErrorValues, errorsArr := checkValues(errorsArr, userParams)
	if hasErrorKeys || hasErrorValues {
		c.JSON(400, gin.H{
			"errors": errorsArr,
		})
		c.Abort()
	}

	// if c.Request.RequestURI == "/personal/" {
	// userParams := userParams.(models.Personal)
	// } else if c.Request.RequestURI == "/aluno/" {
	// userParams := userParams.(models.Aluno)
	// } else {
	// userParams := userParams.(models.User)
	// }

	c.Set("UserParams", userParams)
	c.Next()
}

func checkKeys(errorsArr []string, userParams models.User) (bool, []string) {
	if userParams.Name == "" {
		errorsArr = append(errorsArr, "name is required")
	}

	if userParams.Password == "" {
		errorsArr = append(errorsArr, "password is required")
	}

	if userParams.Nickname == "" {
		errorsArr = append(errorsArr, "nickname is required")
	}

	if userParams.Email == "" {
		errorsArr = append(errorsArr, "email is required")
	}

	if userParams.Cellphone == "" {
		errorsArr = append(errorsArr, "cellphone is required")
	}

	return len(errorsArr) > 0, errorsArr

}

func checkValues(errorsArr []string, userParams models.User) (bool, []string) {
	if len(strings.Trim(userParams.Name, " ")) < 5 {
		errorsArr = append(errorsArr, "Name must have at least 5 caracters")
	}

	if len(strings.Trim(userParams.Password, " ")) < 8 {
		errorsArr = append(errorsArr, "Password must have at lest 8 caracters")
	}

	if len(strings.Trim(userParams.Nickname, " ")) < 5 {
		errorsArr = append(errorsArr, "Nickname must have at least 5 caracters")
	}

	if !validEmail(userParams.Email) {
		errorsArr = append(errorsArr, "Email not valid")
	}

	if strings.Trim(userParams.Cellphone, " ") == "" {
		errorsArr = append(errorsArr, "Cellphone invalid")
	}

	if userParams.Id > 0 {
		errorsArr = append(errorsArr, "To update operations you must use PUT method!")
	}

	return len(errorsArr) > 0, errorsArr

}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
