package gateways

import (
	"github.com/gin-gonic/gin"
	"github.com/halprin/email-conceal/forwarder/context"
	"log"
	"net/http"
)

type RestApiRequestBody struct {
	Url string `json:"url" binding:"required"`
}

func RestApiForwardEmail(arguments map[string]interface{}, applicationContext context.ApplicationContext) error {
	ginContext := arguments["context"].(*gin.Context)

	var json RestApiRequestBody
	err := ginContext.BindJSON(&json)
	if err != nil {
		log.Printf("Failed to parse JSON, error = %+v\n", err)
		ginContext.String(http.StatusBadRequest, "Failed to parse JSON, you must provide a 'url' property set to the URL to the e-mail")
		return err
	}

	log.Println("URL to read e-mail from =", json.Url)

	err = applicationContext.ForwardEmailUsecase(json.Url)
	if err != nil {
		log.Printf("Unable to forward e-mail due to error, %+v\n", err)
		ginContext.String(http.StatusInternalServerError, "E-mail did not forward.  Reach out to the administrator.")
		return err
	}

	log.Println("Forwarded e-mail")
	ginContext.String(http.StatusCreated, "E-mail forwarded successfully")
	return nil
}
