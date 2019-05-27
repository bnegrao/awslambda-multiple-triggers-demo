package main

import (
	"zica-apigateway-lambda/sum"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(DirectInvokeHandler)
}

func DirectInvokeHandler(operands sum.Operands) (result int, err error) {
	return sum.Sum(operands), nil
}
