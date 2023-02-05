package errors

import "github.com/go-sql-driver/mysql"

func CheckError(err *mysql.MySQLError) (bool, string) {
	if err.Number == 1062 {
		return true, "Duplicate data"
	}

	return false, ""
}
