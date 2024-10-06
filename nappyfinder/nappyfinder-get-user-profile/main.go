package main

import (
	"context"
	"errors"
	"nappyfinder/lib/clients"
	"nappyfinder/lib/data"
	"nappyfinder/lib/util"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

var (
	nappyFinderRepository data.NappyFinderRepository
	logger                *logrus.Logger
	isLocal               bool
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger.Debug("handler for nappyfinder-get-user-profile")
	if p, ok := request.PathParameters["userProfileId"]; ok {
		userProfileId := p
		userProfile, err := nappyFinderRepository.GetUserProfileByUserID(ctx, userProfileId)

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

	nappyFinderRepository = &data.NappyFinderDao{
		DB:     dbClient,
		Logger: logger,
	}
	logger.Info("GET - user profile initialized completed ")
}
