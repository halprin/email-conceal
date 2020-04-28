package external

import (
	"github.com/gin-gonic/gin"
)

var restApiApplicationContext = &RestApiApplicationContext{}

func RestApi() {
	router := gin.Default()

	router.POST("/v1/forward", forwardEmail)

	router.Run(":8000")
}

func forwardEmail(context *gin.Context) {
	arguments := map[string]interface{} {
		"context": context,
	}

	_ = restApiApplicationContext.ForwardEmailController(arguments)
}
