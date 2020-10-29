package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"

	"fmt"
	"os"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := cloudwatchlogs.New(sess)

	// Get up to the last 100 log events for LOG-STREAM-NAME
	// in LOG-GROUP-NAME:
	resp, err := svc.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
		Limit:         aws.Int64(100),
		LogGroupName:  aws.String("LOG-GROUP-NAME"),
		LogStreamName: aws.String("LOG-STREAM-NAME"),
	})
	if err != nil {
		fmt.Println("Got error getting log events:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Event messages for stream LOG-STREAM-NAME in log group LOG-GROUP-NAME:")

	gotToken := ""
	nextToken := ""

	for _, event := range resp.Events {
		gotToken = nextToken
		nextToken = *resp.NextForwardToken

		if gotToken == nextToken {
			break
		}

		fmt.Println("  ", *event.Message)
	}
}
