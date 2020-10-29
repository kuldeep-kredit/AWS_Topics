package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

// SubscribeTopic subscribes a user to a topic by their email address
// Inputs:
//     svc is an Amazon SNS service client
//     email is the email address of the user
//     topicARN is the ARN of the topic
// Output:
//     If success, information about the subscription and nil
//     Otherwise, nil and an error from the call to Subscribe
func SubscribeTopic(svc snsiface.SNSAPI, email, topicARN *string) (*sns.SubscribeOutput, error) {
	result, err := svc.Subscribe(&sns.SubscribeInput{
		Endpoint:              email,
		Protocol:              aws.String("email"),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              topicARN,
	})

	return result, err
}

func main() {
	email := flag.String("e", "abhisheksunkariya@gmail.com", "The email address of the user subscribing to the topic")
	topicARN := flag.String("t", "", "The ARN of the topic to which the user subscribes")

	flag.Parse()

	if *email == "" || *topicARN == "" {
		fmt.Println("You must supply an email address and topic ARN")
		fmt.Println("-e EMAIL -t TOPIC-ARN")
		return
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := SubscribeTopic(svc, email, topicARN)
	if err != nil {
		fmt.Println("Got an error subscribing to the topic:")
		fmt.Println(err)
		return
	}

	fmt.Println(*result.SubscriptionArn)
}
