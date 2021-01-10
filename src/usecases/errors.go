package usecases

import "fmt"

type ConcealEmailNotExistError struct {
	ConcealEmailId string
}

func (c ConcealEmailNotExistError) Error() string {
	return fmt.Sprintf("The conceal e-mail %s doesn't exist", c.ConcealEmailId)
}
