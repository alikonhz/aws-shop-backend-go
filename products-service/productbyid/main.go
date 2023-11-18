package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	common "productscommon"
	"slices"
)

func HandleRequest(ctx context.Context, r *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	productID, ok := r.PathParameters["productId"]
	if !ok {
		return buildApiErrorResponse(400, common.ApiError{Message: "productId is required"})
	}

	products := common.ListProducts()
	i := slices.IndexFunc(products, func(p common.Product) bool {
		return p.ID == productID
	})

	if i == -1 {
		return buildApiErrorResponse(404, common.ApiError{Message: fmt.Sprintf("productId %s doesn't exist", productID)})
	}

	return buildOkResponse(products[i])
}

func buildOkResponse(body any) (events.APIGatewayProxyResponse, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(j),
	}, nil
}

func buildApiErrorResponse(statusCode int, apiError common.ApiError) (events.APIGatewayProxyResponse, error) {
	j, err := json.Marshal(apiError)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(j),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
