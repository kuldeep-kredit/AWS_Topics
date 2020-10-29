package main

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

// Define a mock struct to use in unit tests
type mockSNSClient struct {
	snsiface.SNSAPI
}

func (m *mockSNSClient) CreateTopic(input *sns.CreateTopicInput) (*sns.CreateTopicOutput, error) {
	// Check that required inputs exist

	resp := sns.CreateTopicOutput{
		TopicArn: aws.String(""),
	}
	return &resp, nil
}

func TestCreateTopic(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
	t.Log("Starting unit test at " + nowString)

	// mock resource
	topic := ""

	mockSvc := &mockSNSClient{}

	results, err := MakeTopic(mockSvc, &topic)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Created topic with ARN " + *results.TopicArn)
}
