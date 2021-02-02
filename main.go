package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/guardduty"
)

var (
	err error
)

func lambdaHandler(event events.CloudWatchEvent) error {

	if os.Getenv("SLACK_TOKEN") == "" {
		fmt.Println("SLACK_TOKEN env var is empty. Exiting")
		os.Exit(1)
	}

	if os.Getenv("SLACK_CHANNEL_ID") == "" {
		fmt.Println("SLACK_CHANNEL_ID env var is empty. Exiting")
		os.Exit(1)
	}

	guardDutyFindings := guardduty.Finding{}

	err = json.Unmarshal(event.Detail, &guardDutyFindings)

	if err != nil {
		fmt.Printf("Error unmarshalling CloudWatchEvent.Detail %s", err.Error())
		return err
	}

	SendSlackMessage(guardDutyFindings)

	return nil

}

func main() {
	lambda.Start(lambdaHandler)
}
