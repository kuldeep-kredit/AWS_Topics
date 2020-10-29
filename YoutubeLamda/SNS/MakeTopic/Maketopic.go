package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

func MakeTopic(svc snsiface.SNSAPI, topic *string) (*sns.CreateTopicOutput, error) {
	results, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: topic,
	})

	return results, err
}

func main() {
	topic := flag.String("t", "SNSRequest", "The name of the topic")
	flag.Parse()

	if *topic == "" {
		fmt.Println("You must supply the name of the topic")
		fmt.Println("-t TOPIC")
		return
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	results, err := MakeTopic(svc, topic)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(*results.TopicArn)
}
