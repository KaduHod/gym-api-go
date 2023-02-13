package validate

import (
	"api/app/helpers/requests"
	"api/app/models"
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

	hasError, keysErrors := checkUserParams.checkKeys()

	if hasError {
		c.JSON(400, gin.H{
			"errors": keysErrors,
		})

		c.Abort()
	}

	c.Next()
}

func (c *CheckUserPostParams) checkKeys() (bool, *[]string) {
	requiredParams := []string{}

	if c.UserParams.Name == "" {
		requiredParams = append(requiredParams, "name is required")
	}

	if c.UserParams.Password == "" {
		requiredParams = append(requiredParams, "password is required")
	}

	if c.UserParams.Nickname == "" {
		requiredParams = append(requiredParams, "nickname is required")
	}

	if c.UserParams.Email == "" {
		requiredParams = append(requiredParams, "email is required")
	}

	if c.UserParams.Cellphone == "" {
		requiredParams = append(requiredParams, "cellphone is required")
	}

	return len(requiredParams) > 0, &requiredParams

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
