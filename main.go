package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	err error
)

func lambdaHandler(event events.CloudWatchEvent) error {
	guardDutyFindings := GuardDutyFindingDetails{}

	err = json.Unmarshal(event.Detail, &guardDutyFindings)

	if err != nil {
		log.Printf("Error unmarshalling CloudWatchEvent.Detail %s", err.Error())
		return err
	}

	log.Println("GuardDuty event: ", guardDutyFindings)
	SendSlackMessage(guardDutyFindings)

	return nil

}

func main() {
	lambda.Start(lambdaHandler)
}
