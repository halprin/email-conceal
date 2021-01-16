package actualEmail


type ActualEmailUsecase interface {
	Add(actualEmail string) error
}

type ActualEmailUsecaseImpl struct {}

func (receiver ActualEmailUsecaseImpl) Add(actualEmail string) error {
	return nil
}
