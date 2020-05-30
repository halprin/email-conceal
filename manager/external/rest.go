package external

import (
	"github.com/gin-gonic/gin"
)

var applicationContext = &RestApplicationContext{}

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
	httpStatus, jsonMap := applicationContext.ConcealEmailController(genericMap)

	context.JSON(httpStatus, jsonMap)
}

func deleteConcealEmail(context *gin.Context) {
	concealEmailId := context.Param("concealEmailId")

	requestMap := map[string]interface{} {
		"concealEmailId": concealEmailId,
	}

	httpStatus, jsonMap := applicationContext.DeleteConcealEmailController(requestMap)

	context.JSON(httpStatus, jsonMap)
}
