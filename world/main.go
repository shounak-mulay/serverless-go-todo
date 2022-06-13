package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type NameReq struct {
	Name string `json:"name"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error) {

	namereq := NameReq{}
	json.Unmarshal([]byte(req.Body), &namereq)

	body, err := json.Marshal(map[string]interface{}{
		"message": "Okay so your other function also executed successfully!",
		"headers": req.Headers,
		"name":    namereq.Name,
	})

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "world-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
