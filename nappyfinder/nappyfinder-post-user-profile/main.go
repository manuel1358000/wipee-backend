package main

import (
	"context"
	"encoding/json"
	"nappyfinder/lib/clients"
	"nappyfinder/lib/data"
	"nappyfinder/lib/dtos"
	"nappyfinder/lib/models"
	"nappyfinder/lib/util"
	"nappyfinder/lib/validator"
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
	logger.Debug("handler for nappyfinder-post-user-profile")
	userProfileDto := &dtos.UserProfileDto{}
	err := json.Unmarshal([]byte(request.Body), userProfileDto)
	if err != nil {
		return util.NewResponse(logger, request, "Invalid request body", 400, err)
	}

	// Usamos el util de validación
	err = validator.ValidateStruct(userProfileDto)
	if err != nil {
		// Transformamos los errores de validación en un mensaje más legible
		logger.WithFields(logrus.Fields{
			"dto":     userProfileDto,
			"request": request.Body,
		}).Trace("Logging out the dto")
		errorMessage := validator.ParseValidationErrors(err)
		return util.NewResponse(logger, request, errorMessage, 400, err)
	}

	userProfile := models.NewUserProfile(userProfileDto)

	existUserProfile, err := nappyFinderRepository.GetUserProfileDuplicated(ctx, userProfile)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"err":                err,
			"User Profile Exist": existUserProfile,
			"dto":                userProfile,
			"request":            request.Body,
		}).Trace("Logging out the dto")
		return util.NewResponse(logger, request, "", 500, err)
	}
	if len(*existUserProfile) > 0 {
		logger.WithFields(logrus.Fields{
			"User Profile Exist": existUserProfile,
			"dto":                userProfile,
			"request":            request.Body,
			"err":                err,
		}).Error("User Profile duplicated")
		return util.NewResponse(logger, request, "User Profile duplicated", 400, err)
	}

	err = nappyFinderRepository.UpsertUserProfile(ctx, userProfile)

	if err != nil {
		return util.NewResponse(logger, request, "", 500, err)
	}

	logger.WithFields(logrus.Fields{
		"User Profile Exist": existUserProfile,
		"dto":                userProfileDto,
		"request":            request.Body,
		"userProfileEmail":   userProfile.UserProfileEmail,
	}).Trace("Create new Allergy")

	userProfileResponse := userProfile.UserProfileToDto()

	return util.NewResponse(logger, request, userProfileResponse, 200, nil)
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
	logger.Info("POST - user profile initialized completed ")
}
