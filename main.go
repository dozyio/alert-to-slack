package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SlackMessage struct {
	Text string `json:"text"`
}

func handleRequest(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {

		var cloudwatchAlarm events.CloudWatchAlarmSNSPayload

		err := json.Unmarshal([]byte(record.SNS.Message), &cloudwatchAlarm)
		if err != nil {
			log.Print("Couldn't process SNS")
			return
		}

		log.Printf("%s %s %s", cloudwatchAlarm.AlarmName, cloudwatchAlarm.NewStateValue, cloudwatchAlarm.NewStateReason)

		err = toSlack(slackMessage(cloudwatchAlarm))
		if err != nil {
			log.Print("Couldn't send Slack message")
		}
	}
}

func slackMessage(cloudwatchAlarm events.CloudWatchAlarmSNSPayload) SlackMessage {
	return SlackMessage{
		Text: fmt.Sprintf("%s %s %s", cloudwatchAlarm.AlarmName, cloudwatchAlarm.NewStateValue, cloudwatchAlarm.NewStateReason),
	}
}

func toSlack(message SlackMessage) error {
	client := &http.Client{}
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", os.Getenv("SLACK_WEBHOOK"), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return err
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
