package data

import (
	"context"
	"fmt"
	"nappyfinder/lib/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	log "github.com/sirupsen/logrus"
)

type NappyFinderRepository interface {
	GetUserProfileByUserID(ctx context.Context, userProfileId string) (*models.UserProfile, error)
	PostUserProfile(ctx context.Context, userProfile *models.UserProfile) error
	GetUserProfileDuplicated(ctx context.Context, userProfile *models.UserProfile) (*models.ListOfUsersProfile, error)
	UpsertUserProfile(ctx context.Context, userProfile *models.UserProfile) error
}

type NappyFinderDao struct {
	DB     dynamodbiface.DynamoDBAPI
	Logger *log.Logger
}

func (m *NappyFinderDao) GetUserProfileByUserID(ctx context.Context, userProfileId string) (*models.UserProfile, error) {

	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":PK": {
				S: aws.String(fmt.Sprintf("USER#%s", userProfileId)),
			},
			":SK": {
				S: aws.String("#PROFILE"),
			},
		},
		KeyConditionExpression: aws.String("PK = :PK AND begins_with(SK, :SK)"),
		TableName:              aws.String("UserProfileTable"), // Updated table name
		Limit:                  aws.Int64(1),                   // Limit to 1 item
	}

	m.Logger.WithFields(log.Fields{
		"input": input,
	}).Debug("Pre user profile query")

	out, err := m.DB.QueryWithContext(ctx, input)
	if err != nil {
		m.Logger.WithFields(log.Fields{
			"err": err,
		}).Error("Error retrieving user profile")
		return nil, err
	}

	if len(out.Items) == 0 {
		m.Logger.WithFields(log.Fields{}).Info("No user profile found")
		return nil, nil // Or return an error indicating no profile found
	}

	if len(out.Items) > 1 {
		m.Logger.WithFields(log.Fields{}).Warning("Multiple user profiles found for user ID")
	}

	userProfile := &models.UserProfile{}
	err = dynamodbattribute.UnmarshalMap(out.Items[0], userProfile) // Unmarshal the first item
	if err != nil {
		m.Logger.WithFields(log.Fields{
			"err": err,
		}).Error("Error unmarshalling user profile")
		return nil, err
	}

	return userProfile, nil
}

func (m *NappyFinderDao) PostUserProfile(ctx context.Context, userProfile *models.UserProfile) error {
	return nil
}

func (m *NappyFinderDao) GetUserProfileDuplicated(ctx context.Context, userProfile *models.UserProfile) (*models.ListOfUsersProfile, error) {

	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":PK": {
				S: aws.String(userProfile.PK),
			},
			":SK": {
				S: aws.String(userProfile.SK),
			},
		},
		KeyConditionExpression: aws.String("PK = :PK AND begins_with(SK, :SK)"),
		TableName:              aws.String("UserProfileTable"),
	}

	m.Logger.WithFields(log.Fields{
		"input": input,
	}).Debug("User Profile duplicated query")

	out, err := m.DB.QueryWithContext(ctx, input)
	if err != nil {
		m.Logger.WithFields(log.Fields{
			"err": err,
		}).Error("Error retrieving User Profile duplicated")
		return nil, err
	}

	userProfileList := &models.ListOfUsersProfile{}
	if out.Items != nil {
		err = dynamodbattribute.UnmarshalListOfMaps(out.Items, userProfileList)
		if err != nil {
			m.Logger.WithFields(log.Fields{
				"err": err,
			}).Error("Error unmarshalling userProfile duplicated")
			return nil, err
		}
	}

	return userProfileList, nil
}

func (m *NappyFinderDao) UpsertUserProfile(ctx context.Context, userProfile *models.UserProfile) error {

	marshalledEvent, err := dynamodbattribute.MarshalMap(userProfile)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      marshalledEvent,
		TableName: aws.String("UserProfileTable"),
	}

	m.Logger.WithFields(log.Fields{
		"input": input,
	}).Debug("Pre save")

	_, err = m.DB.PutItemWithContext(ctx, input)
	if err != nil {
		m.Logger.WithFields(log.Fields{
			"err": err,
		}).Error("error in saving")
		return err
	}
	return nil
}
