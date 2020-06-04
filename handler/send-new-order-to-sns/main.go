package main

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	message, err := getMessage(request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	err = sendSns(message)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "ok",
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func getMessage(request events.APIGatewayProxyRequest) (string, error) {
	type SendSnsRequest struct {
		Message string `json:"message"`
	}

	sendSnsRequest := SendSnsRequest{
		Message: "",
	}

	err := json.Unmarshal([]byte(request.Body), &sendSnsRequest)

	if err == nil && len(sendSnsRequest.Message) == 0 {
		return sendSnsRequest.Message, errors.New("message should not be empty")
	}

	return sendSnsRequest.Message, err
}

func sendSns(message string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := sns.New(sess)

	input := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(os.Getenv("SNS_ARN")),
	}
	_, err := client.Publish(input)
	return err
}

func main() {
	lambda.Start(Handler)
}
