package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/controllers/actualEmail"
	"github.com/halprin/email-conceal/src/controllers/concealEmail"
)

var applicationContext = context.ApplicationContext{}

func Rest() {
	router := gin.Default()

	v1 := router.Group("/v1")

	//conceal e-mail
	v1.POST("/concealEmail", createConcealEmail)
	v1.PUT("/concealEmail/:concealEmailId", updateConcealEmail)
	v1.DELETE("/concealEmail/:concealEmailId", deleteConcealEmail)

	//actual e-mail
	v1.POST("/actualEmail", createActualEmail)

	_ = router.Run(":8000")
}

func createConcealEmail(context *gin.Context) {
	var genericMap map[string]interface{}

	err := context.BindJSON(&genericMap)
	if err != nil {
		return
	}

	var concealEmailController concealEmail.ConcealEmailController
	applicationContext.Resolve(&concealEmailController)
	httpStatus, jsonMap := concealEmailController.Add(genericMap)

	context.JSON(httpStatus, jsonMap)
}

func deleteConcealEmail(context *gin.Context) {
	concealEmailId := context.Param("concealEmailId")

	requestMap := map[string]interface{} {
		"concealEmailId": concealEmailId,
	}

	var concealEmailController concealEmail.ConcealEmailController
	applicationContext.Resolve(&concealEmailController)
	httpStatus, jsonMap := concealEmailController.Delete(requestMap)

	context.JSON(httpStatus, jsonMap)
}

func updateConcealEmail(context *gin.Context) {
	var requestMap map[string]interface{}
	err := context.BindJSON(&requestMap)
	if err != nil {
		return
	}

	requestMap["concealEmailId"] = context.Param("concealEmailId")

	var concealEmailController concealEmail.ConcealEmailController
	applicationContext.Resolve(&concealEmailController)
	httpStatus, jsonMap := concealEmailController.Update(requestMap)

	context.JSON(httpStatus, jsonMap)
}

func createActualEmail(context *gin.Context) {
	var genericMap map[string]interface{}

	err := context.BindJSON(&genericMap)
	if err != nil {
		return
	}

	var actualEmailController actualEmail.ActualEmailController
	applicationContext.Resolve(&actualEmailController)
	httpStatus, jsonMap := actualEmailController.Add(genericMap)

	context.JSON(httpStatus, jsonMap)
}
