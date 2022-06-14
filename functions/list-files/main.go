package main

import (
	"bytes"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func Handler(c context.Context, req Request) (Response, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return Response{}, err
	}
	client := s3.NewFromConfig(cfg)

	bucket := "serverless-go-test-bucket"
	object := "test.txt"
	obj, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &object,
	})
	if err != nil {
		return Response{}, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(obj.Body)
	myFileContentAsString := buf.String()

	return Response{
		StatusCode: 200,
		Body:       myFileContentAsString,
	}, nil

}

func main() {
	lambda.Start(Handler)
}
