package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/guardduty"
)

// SendSlackMessage sends a message to Slack
// TODO replace this with the slack-go/slack library: https://github.com/slack-go/slack
func SendSlackMessage(findings guardduty.Finding) error {
	var (
		f []SlackFields
		c string
	)

	if *findings.Severity < 5.0 {
		c = "#439FE0"
	} else if *findings.Severity > 7.0 {
		c = "danger"
	} else {
		c = "warning"
	}

	f = append(f,
		SlackFields{Title: "Description", Value: *findings.Description, Short: false},
		SlackFields{Title: "Action Type", Value: *findings.Service.Action.ActionType, Short: true},
		SlackFields{Title: "Finding Type", Value: *findings.Type, Short: false},
	)

	switch *findings.Resource.ResourceType {
	case "AccessKey":
		f = append(f,
			SlackFields{Title: "Suspect Action", Value: "Access Key usage", Short: true},
			SlackFields{Title: "Access Key", Value: "Access key usage", Short: true},
			SlackFields{Title: "User", Value: *findings.Resource.AccessKeyDetails.UserName, Short: true},
		)
	case "Instance":
		f = append(f,
			SlackFields{Title: "Instance ID", Value: *findings.Resource.InstanceDetails.InstanceId, Short: true},
			SlackFields{Title: "Suspect Action", Value: "Instance behavior", Short: true},
		)
	case "S3Bucket":
		f = append(f,
			SlackFields{Title: "Bucket Name", Value: *findings.Resource.S3BucketDetails[0].Name, Short: true},
			SlackFields{Title: "Type", Value: *findings.Resource.S3BucketDetails[0].Type, Short: true},
		)
	default:
		log.Printf("Unknown action type %s", *findings.Resource.ResourceType)
		f = append(f,
			SlackFields{Title: "Suspect Action", Value: "Undetermined", Short: true},
		)
	}

	//log.Println(f)

	sm := SlackMessage{
		Text: fmt.Sprintf("GuardDuty Alert in %s", *findings.Region),
		Attachments: []SlackAttachments{
			{
				Text:   "New Findings",
				Color:  c,
				Title:  "GuardDuty Alert",
				Fields: f,
			},
		},
	}

	client := &http.Client{Timeout: 4 * time.Second}

	d, err := json.Marshal(sm)

	if err != nil {
		log.Printf("Hit error marshalling Slack JSON - %s", err.Error())
		return err
	}

	req, err := http.NewRequest("POST", os.Getenv("SLACK_WEBHOOK"), bytes.NewBuffer(d))

	if err != nil {
		log.Printf("Hit error creating http request - %s", err.Error())
		return err
	}

	resp, err := client.Do(req)

	rb, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Hit error making request to Slack - %s", err.Error())
		return err
	}

	if resp.StatusCode != 200 {
		log.Printf("Hit error sending Slack request - status code: %d - body: %s", resp.StatusCode, string(rb))
		return errors.New("Error sending Slack request")
	}

	return nil
}
