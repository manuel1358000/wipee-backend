package main

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"wipee/lib/clients"
	"wipee/lib/data"
	"wipee/lib/dtos"
	"wipee/lib/models"
	"wipee/lib/util"
	"wipee/lib/validator"

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
	logger.Debug("handler for post-user-profile")
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

	existUserProfile, err := wipeeRepository.GetUserProfileDuplicated(ctx, userProfile)

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

	err = wipeeRepository.UpsertUserProfile(ctx, userProfile)

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

	wipeeRepository = &data.WipeeDao{
		DB:     dbClient,
		Logger: logger,
	}
	logger.Info("POST - user profile initialized completed ")
}
