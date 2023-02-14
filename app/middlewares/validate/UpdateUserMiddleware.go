package validate

import (
	"api/app/helpers/requests"
	"api/app/models"
	"net/mail"
	"strings"

	"github.com/gin-gonic/gin"
)

type CheckUserPutParams struct {
	UserParams *models.User
}

func UpdateUserMiddleware(c *gin.Context) {
	var userParams models.User
	var errorsArr []string

	requests.GetBodyJson(c.Request.Body, &userParams)

	checkUserParams := CheckUserPutParams{
		UserParams: &userParams,
	}

	hasErrorKeys, errorsArr := checkUserParams.checkKeys(errorsArr)
	hasErrorValues, errorsArr := checkUserParams.checkValues(errorsArr)
	if hasErrorKeys || hasErrorValues {
		c.JSON(400, gin.H{
			"errors": errorsArr,
		})
		c.Abort()
	}

	c.Next()
}

func (c *CheckUserPutParams) checkKeys(errorsArr []string) (bool, []string) {
	if c.UserParams.Id == 0 {
		errorsArr = append(errorsArr, "Id is required")
	}

	return len(errorsArr) > 0, errorsArr

}

func (c *CheckUserPutParams) checkValues(errorsArr []string) (bool, []string) {
	if len(strings.Trim(c.UserParams.Name, " ")) < 5 {
		errorsArr = append(errorsArr, "Name must have at least 5 caracters")
	}

	if len(strings.Trim(c.UserParams.Password, " ")) < 8 {
		errorsArr = append(errorsArr, "Password must have at lest 8 caracters")
	}

	if len(strings.Trim(c.UserParams.Nickname, " ")) < 5 {
		errorsArr = append(errorsArr, "Nickname must have at least 5 caracters")
	}

	if !c.validEmail() {
		errorsArr = append(errorsArr, "Email not valid")
	}

	if strings.Trim(c.UserParams.Cellphone, " ") == "" {
		errorsArr = append(errorsArr, "Cellphone invalid")
	}

	if c.UserParams.Id > 0 {
		errorsArr = append(errorsArr, "To update operations you must use PUT method!")
	}

	return len(errorsArr) > 0, errorsArr

}

func (c *CheckUserPutParams) validEmail() bool {
	_, err := mail.ParseAddress(c.UserParams.Email)
	return err == nil
}
