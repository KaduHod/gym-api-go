package middlewares

import (
	"api/app/helpers/requests"
	"api/app/models"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type CheckUserPostParams struct {
	UserParams *models.User
}

func CreateUserMiddleware(c *gin.Context) {
	var userParams models.User
	requests.GetBodyJson(c.Request.Body, &userParams)
	checkUserParams := CheckUserPostParams{
		UserParams: &userParams,
	}

	fmt.Println(userParams)

	hasError, keysErrors := checkUserParams.checkKeys()

	if hasError {
		c.JSON(400, gin.H{
			"errors": keysErrors,
		})

		return
	} else {
		c.Next()
	}

	fmt.Println(keysErrors)

}

func (c *CheckUserPostParams) checkKeys() (bool, *[]string) {
	requiredParams := []string{}

	if c.UserParams.Name == "" {
		requiredParams = append(requiredParams, "name")
	}

	if c.UserParams.Password == "" {
		requiredParams = append(requiredParams, "password")
	}

	if c.UserParams.Nickname == "" {
		requiredParams = append(requiredParams, "nickname")
	}

	if c.UserParams.Email == "" {
		requiredParams = append(requiredParams, "email")
	}

	if c.UserParams.Cellphone == "" {
		requiredParams = append(requiredParams, "cellphone")
	}

	length := len(requiredParams)

	return length > 0, &requiredParams

}

func (c *CheckUserPostParams) checkValues(msgs *[]string) {
	valuesErrorMsgs := []string{}
	if strings.Trim(c.UserParams.Name, " ") == "" {
		valuesErrorMsgs = append(valuesErrorMsgs, "name")
	}

	if strings.Trim(c.UserParams.Password, " ") == "" {
		valuesErrorMsgs = append(valuesErrorMsgs, "password")
	}

	if strings.Trim(c.UserParams.Nickname, " ") == "" {
		valuesErrorMsgs = append(valuesErrorMsgs, "nickname")
	}

	if strings.Trim(c.UserParams.Email, " ") == "" {
		valuesErrorMsgs = append(valuesErrorMsgs, "email")
	}

	if strings.Trim(c.UserParams.Cellphone, " ") == "" {
		valuesErrorMsgs = append(valuesErrorMsgs, "cellphone")
	}

}
