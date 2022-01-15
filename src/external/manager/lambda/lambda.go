package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/halprin/email-conceal/src/external/manager/localRest"
)

var ginLambda *ginadapter.GinLambda

func PrepareLambda() {
	ginRouter := localRest.RestConfiguration()

	ginLambda = ginadapter.New(ginRouter)
}

func LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
