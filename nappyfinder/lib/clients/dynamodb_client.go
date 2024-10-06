package clients

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// NewDynamoDBClient initializes a DynamoDB session with optional local configuration
func NewDynamoDBClient(isLocal bool) dynamodbiface.DynamoDBAPI {

	c := &aws.Config{
		Region: aws.String("us-west-2")}
	if isLocal {
		c.Endpoint = aws.String("http://docker.for.mac.host.internal:4566")
	}
	sess := session.Must(session.NewSession(c))
	svc := dynamodb.New(sess)
	return dynamodbiface.DynamoDBAPI(svc)
}
