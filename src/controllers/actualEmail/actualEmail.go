package actualEmail

import (
	"github.com/halprin/email-conceal/src/context"
	actualEmailUsecase2 "github.com/halprin/email-conceal/src/usecases/actualEmail"
	"log"
	"net/http"
)

var applicationContext = context.ApplicationContext{}
var actualEmailUsecase actualEmailUsecase2.ActualEmailUsecase

type ActualEmailController struct {}

func (receiver ActualEmailController) Init() {
	applicationContext.Resolve(&actualEmailUsecase)
}

func (receiver ActualEmailController) Add(arguments map[string]interface{}) (int, map[string]string) {
	actualEmail, actualEmailValid := arguments["email"].(string)
	if !actualEmailValid {
		errorString := "E-mail was not supplied or was not a string"
		log.Printf(errorString)
		jsonMap := map[string]string{
			"error": errorString,
		}
		return http.StatusBadRequest, jsonMap
	}

	log.Printf("Parsed out the actual e-mail %s", actualEmail)

	//var actualEmailUsecase actualEmailUsecase2.ActualEmailUsecase
	//applicationContext.Resolve(&actualEmailUsecase)
	err := actualEmailUsecase.Add(actualEmail)

	//TODO: do more error handling based on whatever errors can come back from the usecase
	if err != nil {
		log.Printf("Another error occured, %+v", err)
		jsonMap := map[string]string{
			"error": "An unknown error occurred",
		}
		return http.StatusInternalServerError, jsonMap
	}

	return http.StatusCreated, nil
}
