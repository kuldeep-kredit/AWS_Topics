package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

// GetSubscriptions retrieves a list of your Amazon SNS subscriptions
// Inputs:
//     svc is an Amazon SNS service client
// Output:
//     If success, information about your subscriptions and nil
//     Otherwise, nil and an error from the call to ListSubscriptions
func GetSubscriptions(svc snsiface.SNSAPI) (*sns.ListSubscriptionsOutput, error) {
	result, err := svc.ListSubscriptions(nil)

	return result, err
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := GetSubscriptions(svc)
	if err != nil {
		fmt.Println("Got an error retrieving the subscriptions:")
		fmt.Println(err)
		return
	}

	fmt.Println("Topic ARN")
	fmt.Println("Subscription ARN")
	fmt.Println("-------------------------")
	for _, s := range result.Subscriptions {
		fmt.Println(*s.TopicArn)
		fmt.Println(*s.SubscriptionArn)

		fmt.Println("")
	}
}
