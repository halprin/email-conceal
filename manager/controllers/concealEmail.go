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
	sourceEmail, sourceEmailValid := arguments["email"].(string)
	if !sourceEmailValid {
		errorString := "E-mail was not supplied or was not a string"
		log.Printf(errorString)
		jsonMap := map[string]string{
			"error": errorString,
		}
		return http.StatusBadRequest, jsonMap
	}

	var description *string = nil
	descriptionRaw, descriptionExists := arguments["description"]
	if descriptionExists {
		localDescription, descriptionValid := descriptionRaw.(string)
		description = &localDescription
		if !descriptionValid {
			errorString := "Description was not a string"
			log.Printf(errorString)
			jsonMap := map[string]string{
				"error": errorString,
			}
			return http.StatusBadRequest, jsonMap
		}
	}

	if description != nil {
		log.Printf("E-mail to conceal and description = %s, %s", sourceEmail, *description)
	} else {
		log.Printf("E-mail to conceal with no description = %s", sourceEmail)
	}

	concealedEmail, err := applicationContext.Usecases().AddConcealEmail(sourceEmail, description)

	if errors.Is(err, entities.InvalidEmailAddressError) {
		errorString := fmt.Sprintf("E-mail %s is invalid", sourceEmail)
		log.Printf(errorString)
		jsonMap := map[string]string{
			"error": errorString,
		}
		return http.StatusBadRequest, jsonMap
	} else if errors.Is(err, entities.DescriptionTooShortError) || errors.Is(err, entities.DescriptionTooLongError) {
		errorString := err.Error()
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
