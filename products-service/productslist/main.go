package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	common "productscommon"
)

func HandleRequest(ctx context.Context, r *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	j, err := json.Marshal(common.ListProducts())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(j),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
