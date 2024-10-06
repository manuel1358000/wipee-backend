package models

import (
	"fmt"
	"nappyfinder/lib/dtos"
	"time"

	"github.com/segmentio/ksuid"
)

/*
Type(s)
*/
type UserProfile struct {
	UserProfileId    string     `dynamodbav:"UserProfileId,omitempty"`
	UserProfileName  string     `dynamodbav:"UserProfileName,omitempty"`
	UserProfileEmail string     `dynamodbav:"UserProfileEmail,omitempty"`
	Status           string     `dynamodbav:"Status,omitempty"`
	IsDelete         bool       `dynamodbav:"IsDelete,omitempty"`
	Country          string     `dynamodbav:"Country,omitempty"`
	PathProfilePhoto string     `dynamodbav:"PathProfilePhoto,omitempty"`
	BirthDateTime    *time.Time `dynamodbav:"BirthDateTime,omitempty"`
	CreateDateTime   *time.Time `dynamodbav:"CreateDateTime,omitempty"`
	DeleteDatetime   *time.Time `dynamodbav:"DeleteDateTime,omitempty"`
	ModifyDatetime   *time.Time `dynamodbav:"ModifyDateTime,omitempty"`

	PK string `dynamodbav:"PK"`
	SK string `dynamodbav:"SK"`
}

type ListOfUsersProfile []*UserProfile

func (u *UserProfile) UserProfileToDto() *dtos.UserProfileDto {
	return &dtos.UserProfileDto{
		UserProfileId:    u.UserProfileId,
		UserProfileName:  u.UserProfileName,
		UserProfileEmail: u.UserProfileEmail,
		Status:           u.Status,
		IsDelete:         u.IsDelete,
		Country:          u.Country,
		PathProfilePhoto: u.PathProfilePhoto, // Correct this line as needed (bool vs string)
		BirthDateTime:    u.BirthDateTime,
		CreateDateTime:   u.CreateDateTime,
		DeleteDateTime:   u.DeleteDatetime,
		ModifyDateTime:   u.ModifyDatetime,
	}
}

func NewUserProfile(userProfileDto *dtos.UserProfileDto) *UserProfile {
	id := ksuid.New().String()
	tNow := time.Now()
	a := &UserProfile{
		UserProfileId:    id,
		UserProfileName:  userProfileDto.UserProfileName,
		UserProfileEmail: userProfileDto.UserProfileEmail,
		Status:           "NOT-VERIFIED",
		IsDelete:         false,
		Country:          "GT",
		PathProfilePhoto: userProfileDto.PathProfilePhoto, // Correct this line as needed (bool vs string)
		BirthDateTime:    userProfileDto.BirthDateTime,
		CreateDateTime:   &tNow,
		PK:               NewUserProfilePK(id),
		SK:               NewUserProfileSK(),
	}
	return a
}

func NewUserProfilePK(UserProfileId string) string {
	return fmt.Sprintf("USER#%s", UserProfileId)
}

func NewUserProfileSK() string {
	return "#PROFILE"
}
