package account

import (
	"fmt"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/entities"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"log"
	"net/http"
)
import accountUsecase2 "github.com/halprin/email-conceal/src/usecases/account"

var applicationContext = context.ApplicationContext{}
var accountUsecase accountUsecase2.AccountUsecase

type AccountController struct{}

func (receiver AccountController) Init() {
	applicationContext.Resolve(&accountUsecase)
}

func (receiver AccountController) New(arguments map[string]interface{}) (int, map[string]string) {
	username, usernameValid := arguments["username"].(string)

	if !usernameValid {
		errorString := "Username was not supplied or was not a string"
		log.Printf(errorString)
		jsonMap := map[string]string{
			"error": errorString,
		}
		return http.StatusBadRequest, jsonMap
	}

	password, passwordValid := arguments["password"].(string)

	if !passwordValid {
		errorString := "Password was not supplied or was not a string"
		log.Printf(errorString)
		jsonMap := map[string]string{
			"error": errorString,
		}
		return http.StatusBadRequest, jsonMap
	}

	err := accountUsecase.Create(username, password)

	if errors.Is(err, entities.InvalidEmailAddressError) {
		errorString := fmt.Sprintf("Username %s is invalid", username)
		log.Printf(errorString)
		jsonMap := map[string]string{
			"error": errorString,
		}
		return http.StatusBadRequest, jsonMap
	} else if err != nil {
		log.Printf("Another error occured, %+v", err)
		jsonMap := map[string]string{
			"error": "An unknown error occurred",
		}
		return http.StatusInternalServerError, jsonMap
	}

	log.Println("New account =", username)

	jsonMap := map[string]string{
		"account": username,
	}

	return http.StatusCreated, jsonMap
}
