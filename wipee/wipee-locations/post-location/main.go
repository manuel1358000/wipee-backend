package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"wipee/lib/clients"
	"wipee/lib/data"
	"wipee/lib/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

var (
	wipeeRepository data.WipeeRepository
	logger          *logrus.Logger
	isLocal         bool
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger.Debug("handler for post-location")
	fmt.Printf("Request received: %+v\n", request)

	return util.NewResponse(logger, request, "Post Locations", 200, nil)
}

func main() {
	lambda.Start(Handler)
}

func init() {
	isLocal, _ = strconv.ParseBool(os.Getenv("IS_LOCAL"))
	dbClient := clients.NewDynamoDBClient(isLocal)

	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: isLocal,
	})

	wipeeRepository = &data.WipeeDao{
		DB:     dbClient,
		Logger: logger,
	}
	logger.Info("POST - post locations initialized completed ")
}
