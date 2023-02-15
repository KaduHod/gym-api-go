package validate

import (
	"api/app/helpers/requests"
	"api/app/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdateUserMiddleware[T models.UserType](c *gin.Context) {
	var userParams T
	var errorsArr []string

	requests.GetBodyJson(c.Request.Body, &userParams)

	hasErrorKeys, errorsArr := checkUpdateKeys(errorsArr, userParams)
	hasErrorValues, errorsArr := checkUpdateValues(errorsArr, userParams)
	if hasErrorKeys || hasErrorValues {
		c.JSON(400, gin.H{
			"errors": errorsArr,
		})
		c.Abort()
	}
	c.Set("UserParams", userParams)
	c.Next()
}

func checkUpdateKeys[T models.UserType](errorsArr []string, user T) (bool, []string) {
	if user.GetId() == 0 {
		errorsArr = append(errorsArr, "Id is required")
	}

	return len(errorsArr) > 0, errorsArr

}

func checkUpdateValues[T models.UserType](errorsArr []string, user T) (bool, []string) {
	if len(strings.Trim(user.GetName(), " ")) < 5 && strings.Trim(user.GetName(), " ") != "" {
		errorsArr = append(errorsArr, "Name must have at least 5 caracters")
	}

	if len(strings.Trim(user.GetPassword(), " ")) < 8 && strings.Trim(user.GetPassword(), " ") != "" {
		errorsArr = append(errorsArr, "Password must have at lest 8 caracters")
	}

	if len(strings.Trim(user.GetNickname(), " ")) < 5 && strings.Trim(user.GetNickname(), " ") != "" {
		errorsArr = append(errorsArr, "Nickname must have at least 5 caracters")
	}

	if !validEmail(user.GetEmail()) && strings.Trim(user.GetEmail(), " ") != "" {
		errorsArr = append(errorsArr, "Email not valid")
	}

	if strings.Trim(user.GetCellphone(), " ") != "" {
		errorsArr = append(errorsArr, "Cellphone invalid")
	}

	return len(errorsArr) > 0, errorsArr
}
