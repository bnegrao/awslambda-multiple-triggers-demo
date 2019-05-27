package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"zica-apigateway-lambda/sum"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ApiGatewayHandler)
}

func ApiGatewayHandler(ctx context.Context, apiEvent events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	operands, validationError := validateInput(apiEvent.Body)
	if validationError != nil {
		return createError(400, "Validation Error: "+validationError.Error()), nil
	}

	sumResult := sum.Sum(operands)

	return newHttpResponse(200, map[string]string{
		"content-type": "text/plain; charset=UTF-8",
	}, strconv.Itoa(sumResult)), nil
}

func validateInput(requestBody string) (operands sum.Operands, validationError error) {
	operands = sum.Operands{}

	err := json.Unmarshal([]byte(requestBody), &operands)
	if err != nil {
		return operands, err
	}

	return operands, nil
}

func createError(statusCode int, message string) events.APIGatewayProxyResponse {
	msg := map[string]string{"message": message}
	messageJson, _ := json.Marshal(msg)

	fmt.Println("ERROR: ", message)

	return newHttpResponse(statusCode, nil, string(messageJson))
}

func newHttpResponse(statusCode int, headers map[string]string, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}
}
