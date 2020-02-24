package external

import (
	"github.com/gin-gonic/gin"
)

var applicationContext = &RestApiApplicationContext{}

func RestApi() {
	router := gin.Default()

	router.POST("/v1/forward", forwardEmail)

	router.Run(":8080")
}

func forwardEmail(context *gin.Context) {
	arguments := map[string]interface{} {
		"context": context,
	}

	_ = applicationContext.ForwardEmailGateway(arguments)
}
