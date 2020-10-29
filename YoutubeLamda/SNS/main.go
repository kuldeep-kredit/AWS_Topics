package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must supply a topic name")
		fmt.Println("Usage: go run SnsCreateTopic.go TOPIC")
		os.Exit(1)
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(os.Args[1]),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.TopicArn)
}

// func main() {

// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String("us-east-2"),
// 	})

// 	if err != nil {
// 		fmt.Println("NewSession error:", err)
// 		return
// 	}

// 	client := sns.New(sess)
// 	input := &sns.PublishInput{
// 		Message:  aws.String("Hello world!"),
// 		TopicArn: aws.String("arn:aws:sns:us-east-2:313049956234:SNSRequest"),
// 	}

// 	result, err := client.Publish(input)
// 	if err != nil {
// 		fmt.Println("Publish error:", err)
// 		return
// 	}

// 	fmt.Println("your message", result)
// }
