# guardduty-slack

## Slack Setup

Create and install a Slack application. Make sure your app has `chat:write` permission to the correct channel(s) in the "OAuth & Permissions" section of the app configuration. `/invite` the bot user to the channel you want messages sent to.

Grab the app's "Bot User OAuth Access Token" and the channel ID of the channel you want to send messages to. 

This is not compatible with legacy Slack incoming webhooks.

## Terraform

Check the `terraform/` directory for an example of how to set this up.