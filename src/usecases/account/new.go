package account

type Account struct {
}

func (receiver Account) Create(emailUsername string, password string) error {
	//TODO: write emailUsername and hashed/salted version of password to database
	//TODO: call actualEmail usecase

	return nil
}
