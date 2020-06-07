package controllers

import (
	"fmt"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/entities"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
	"log"
	"net/http"
)


func HttpConcealEmailController(arguments map[string]interface{}, applicationContext context.ApplicationContext) (int, map[string]string) {
	sourceEmail, valid := arguments["email"].(string)
	if !valid {
		errorString := "E-mail was not supplied or was not a string"
		log.Printf(errorString)
		jsonMap := map[string]string{
			"error": errorString,
		}
		return http.StatusBadRequest, jsonMap
	}

	log.Println("E-mail to conceal =", sourceEmail)

	concealedEmail, err := applicationContext.Usecases().AddConcealEmail(sourceEmail)

	if errors.Is(err, entities.InvalidEmailAddressError) {
		errorString := fmt.Sprintf("E-mail %s is invalid", sourceEmail)
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

	log.Println("Concealed e-mail address =", concealedEmail)

	jsonMap := map[string]string{
		"concealedEmail": concealedEmail,
	}

	return http.StatusCreated, jsonMap
}

func HttpDeleteConcealEmailController(arguments map[string]interface{}, applicationContext context.ApplicationContext) (int, map[string]string) {
	concealEmailId, valid := arguments["concealEmailId"].(string)

	if !valid {
		errorString := "Conceal e-mail ID was not supplied or was not a string"
		log.Printf(errorString)
		jsonMap := map[string]string{
			"error": errorString,
		}
		return http.StatusBadRequest, jsonMap
	}

	log.Println("Conceal E-mail ID to delete =", concealEmailId)

	err := applicationContext.Usecases().DeleteConcealEmail(concealEmailId)
	if err != nil {
		log.Printf("Some error occured while trying to delete the conceal e-mail, %+v", err)
		jsonMap := map[string]string{
			"error": "An unknown error occurred",
		}
		return http.StatusInternalServerError, jsonMap
	}

	jsonMap := map[string]string{}
	return http.StatusNoContent, jsonMap
}
