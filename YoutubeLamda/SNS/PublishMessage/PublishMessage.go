package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

// PublishMessage publishes a message to an Amazon SNS topic
// Inputs:
//     svc is an Amazon SNS service object
//     msg is the message to publish
//     topicARN is the Amazon Resource Name (ARN) of the topic to publish through
// Output:
//     If success, information about the publication and nil
//     Otherwise, nil and an error from the call to Publish
func PublishMessage(svc snsiface.SNSAPI, msg, topicARN *string) (*sns.PublishOutput, error) {
	result, err := svc.Publish(&sns.PublishInput{
		Message:  msg,
		TopicArn: topicARN,
	})

	return result, err
}

func main() {
	msg := flag.String("m", "somethingwent wrong", "The message to send to the subscribed users of the topic")
	topicARN := flag.String("t", "", "The ARN of the topic to which the user subscribes")

	flag.Parse()

	if *msg == "" || *topicARN == "" {
		fmt.Println("You must supply a message and topic ARN")
		fmt.Println("-m MESSAGE -t TOPIC-ARN")
		return
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := PublishMessage(svc, msg, topicARN)
	if err != nil {
		fmt.Println("Got an error publishing the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Message ID: ", result)
}
