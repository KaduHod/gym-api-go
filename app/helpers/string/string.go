package string

import "strconv"

func ToInteger(value string) (int, error) {
	integer, err := strconv.Atoi(value)
	return integer, err
}
