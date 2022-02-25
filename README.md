# Alert To Slack

AWS Cloudwatch alerts to slack

# Setup

* Create a AWS Lambda called alert-to-slack
* Set handler to main
* Add environment variable "SLACK_WEBHOOK" to your slack endpoint for incoming webhooks
* Create SNS topic with the new Lambda as a subscriber
* Create cloudwatch alarms that trigger the new SNS topic

# Build and deploy
```
make all
```
