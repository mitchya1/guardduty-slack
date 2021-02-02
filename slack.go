package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/slack-go/slack"
)

// SendSlackMessage sends a message to Slack
func SendSlackMessage(findings guardduty.Finding) error {
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	var (
		f []slack.AttachmentField
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
		slack.AttachmentField{Title: "Description", Value: *findings.Description, Short: false},
		slack.AttachmentField{Title: "Action Type", Value: *findings.Service.Action.ActionType, Short: true},
		slack.AttachmentField{Title: "Finding Type", Value: *findings.Type, Short: false},
	)

	switch *findings.Resource.ResourceType {
	case "AccessKey":
		var ak string
		if findings.Resource.AccessKeyDetails.AccessKeyId != nil {
			ak = *findings.Resource.AccessKeyDetails.AccessKeyId
		} else {
			ak = "Access key usage"
		}
		f = append(f,
			slack.AttachmentField{Title: "Suspect Action", Value: "Access Key usage", Short: true},
			slack.AttachmentField{Title: "Access Key", Value: ak, Short: true},
			slack.AttachmentField{Title: "User", Value: *findings.Resource.AccessKeyDetails.UserName, Short: true},
		)
	case "Instance":
		f = append(f,
			slack.AttachmentField{Title: "Instance ID", Value: *findings.Resource.InstanceDetails.InstanceId, Short: true},
			slack.AttachmentField{Title: "Suspect Action", Value: "Instance behavior", Short: true},
		)
	case "S3Bucket":
		f = append(f,
			slack.AttachmentField{Title: "Bucket Name", Value: *findings.Resource.S3BucketDetails[0].Name, Short: true},
			slack.AttachmentField{Title: "Type", Value: *findings.Resource.S3BucketDetails[0].Type, Short: true},
		)
	default:
		log.Printf("Unknown action type %s", *findings.Resource.ResourceType)
		f = append(f,
			slack.AttachmentField{Title: "Suspect Action", Value: "Undetermined", Short: true},
			slack.AttachmentField{Title: "Resource", Value: *findings.Resource.ResourceType},
		)
	}

	sm := slack.Attachment{
		Pretext: fmt.Sprintf("New Findings in %s", *findings.Region),
		Color:   c,
		Title:   "GuardDuty Alert",
		Fields:  f,
	}

	if _, _, err = api.PostMessage(
		os.Getenv("SLACK_CHANNEL_ID"),
		slack.MsgOptionAttachments(sm),
		slack.MsgOptionAsUser(true),
	); err != nil {
		fmt.Printf("Error sending slack message - %s", err.Error())
		return err
	}

	return nil
}
