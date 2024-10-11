package main

import (
	"context"
	"errors"
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
	logger.Debug("handler for get-user-profile")
	if p, ok := request.PathParameters["userProfileId"]; ok {
		userProfileId := p
		userProfile, err := wipeeRepository.GetUserProfileByUserID(ctx, userProfileId)

		if err != nil {
			return util.NewResponse(logger, request, "Error", 500, errors.New("internal server error"))
		}

		userProfileDto := userProfile.UserProfileToDto()

		return util.NewResponse(logger, request, userProfileDto, 200, nil)
	}

	return util.NewResponse(logger, request, "POST", 400, errors.New("missing path parameter"))
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
	logger.Info("GET - user profile initialized completed ")
}
