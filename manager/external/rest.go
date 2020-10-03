package external

import (
	"github.com/gin-gonic/gin"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/controllers"
	"github.com/halprin/email-conceal/manager/external/restContext"
)

var applicationContext = context.ApplicationContext{}
var concealEmailController controllers.ConcealEmailController

func init() {
	restContext.Init()
	applicationContext.Resolve(&concealEmailController)
}

func Rest() {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.POST("/concealEmail", createConcealEmail)
	v1.DELETE("/concealEmail/:concealEmailId", deleteConcealEmail)

	_ = router.Run(":8000")
}

func createConcealEmail(context *gin.Context) {
	var genericMap map[string]interface{}

	err := context.BindJSON(&genericMap)
	if err != nil {
		return
	}
	httpStatus, jsonMap := concealEmailController.Add(genericMap)

	context.JSON(httpStatus, jsonMap)
}

func deleteConcealEmail(context *gin.Context) {
	concealEmailId := context.Param("concealEmailId")

	requestMap := map[string]interface{} {
		"concealEmailId": concealEmailId,
	}

	httpStatus, jsonMap := concealEmailController.Delete(requestMap)

	context.JSON(httpStatus, jsonMap)
}
